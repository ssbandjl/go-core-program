package controller

import (
	"dms-backend/database"
	"dms-backend/job"
	"dms-backend/model"
	"dms-backend/runtime"
	"dms-backend/util"
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"strconv"
	"strings"
	"time"

	batchv1 "k8s.io/api/batch/v1"
	"k8s.io/api/batch/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func MysqlInstanceList(c *gin.Context) {
	pageSizeDefault, _ := runtime.Cfg.GetValue("utils", "PAGE_SIZE")

	search := c.DefaultQuery("search", "")
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", pageSizeDefault))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

	clusterID := c.DefaultQuery("cluster_id", "")
	//fmt.Printf("MysqlInstanceList clusterID:%v\n", clusterID)

	var total int
	var instances []model.MysqlInstance
	countDB := database.DB.Order("id desc").Model(&model.MysqlInstance{}).Where("is_delete = ?", false)
	dataDB := database.DB.Order("id desc").Offset((page-1)*pageSize).Limit(pageSize).Preload("User").Preload("KubernetesInfo").Where("is_delete = ?", false)
	//fmt.Printf("MysqlInstanceList countDB:%v\n", countDB)
	//fmt.Printf("MysqlInstanceList dataDB:%v\n", dataDB)

	if clusterID != "" {
		countDB = countDB.Where("kubernetes_id = ?", clusterID)
		dataDB = dataDB.Where("kubernetes_id = ?", clusterID)
	}
	if search != "" {
		countDB = countDB.Where("name like ? or alias like ?", fmt.Sprintf("%%%s%%", search), fmt.Sprintf("%%%s%%", search))
		dataDB = dataDB.Where("name like ? or alias like ?", fmt.Sprintf("%%%s%%", search), fmt.Sprintf("%%%s%%", search))
	}

	countDB.Count(&total)
	dataDB.Find(&instances)

	for _, v := range instances {
		fmt.Printf("\nMysqlInstanceList 获取实例列表:\n%+v\n", v)
	}
	c.AbortWithStatusJSON(200, gin.H{"success": true, "total": total, "data": instances, "page": page, "page_size": pageSize})
	return
}

func MysqlInstanceCreate(c *gin.Context) {
	postDataByte, _ := ioutil.ReadAll(c.Request.Body)
	var postData map[string]interface{}
	err := json.Unmarshal(postDataByte, &postData)
	//fmt.Printf("MysqlInstanceCreate postData:\n%+v\n", postData)
	//[alias:test instance_type:single kubernetes_controller:nginx kubernetes_controller_type:Deployment kubernetes_id:1 kubernetes_ns:cm-ota port:111 root_password:root_password]
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"success": false, "message": err.Error()})
		return
	}

	kubernetesController := postData["kubernetes_controller"].(string)
	kubernetesControllerType := postData["kubernetes_controller_type"].(string)
	kubernetesNS := postData["kubernetes_ns"].(string)
	kubernetesID := int(postData["kubernetes_id"].(float64))
	port := int(postData["port"].(float64))
	alias := postData["alias"].(string)
	rootPassword := postData["root_password"].(string)
	instanceType := postData["instance_type"].(string)

	instance := model.MysqlInstance{
		Name:                     fmt.Sprintf("%s-%s", kubernetesNS, kubernetesController),
		KubernetesNS:             kubernetesNS,
		KubernetesInfoID:         kubernetesID,
		Port:                     port,
		KubernetesControllerType: kubernetesControllerType,
		KubernetesController:     kubernetesController,
		Alias:                    alias,
		RootPassword:             b64.StdEncoding.EncodeToString([]byte(rootPassword)),
		IsDelete:                 false,
		CreateAt:                 time.Now(),
		InstanceType:             instanceType,
		UserID:                   sessions.Default(c).Get("user").(model.User).ID,
	}

	var kubernetesInfo model.KubernetesInfo
	database.DB.Where("id = ?", kubernetesID).First(&kubernetesInfo)
	if kubernetesInfo.ID == 0 {
		c.AbortWithStatusJSON(200, gin.H{"success": false, "message": "获取集群信息失败"})
		return
	}

	var isExistInstance model.MysqlInstance
	database.DB.Where("name = ? and kubernetes_id = ? and is_delete = ?", instance.Name, kubernetesID, false).First(&isExistInstance)
	if isExistInstance.ID > 0 {
		c.AbortWithStatusJSON(200, gin.H{"success": false, "message": "实例已存在"})
		return
	}

	client := util.KubernetesClient{Host: kubernetesInfo.ClusterURL, Port: kubernetesInfo.Port, Token: kubernetesInfo.Token, Scheme: "https"}
	// 获取svc
	svcs := client.GetSvcs(instance.KubernetesNS, v1.ListOptions{})
	if svcs == nil {
		c.AbortWithStatusJSON(200, gin.H{"success": false, "message": "创建实例，获取对应svc失败"})
		return
	}

	for _, svc := range svcs.Items {
		if strings.Contains(instance.KubernetesController, svc.ObjectMeta.Name) {
			for _, port := range svc.Spec.Ports {
				if port.Port == int32(instance.Port) {
					instance.KubernetesSVC = svc.ObjectMeta.Name
					break
				}
			}
		}
	}
	if instance.KubernetesSVC == "" {
		c.AbortWithStatusJSON(200, gin.H{"success": false, "message": "创建实例，找不到对应svc"})
		return
	}

	_, agentPort := util.GetAgent(kubernetesInfo.Namespace, client)
	dt, err := util.HttpRequest(fmt.Sprintf("http://%s:%d/api/mysql/connect/check?host=%s&port=%d&username=root&password=%s", kubernetesInfo.ClusterURL, agentPort, fmt.Sprintf("%s.%s", instance.KubernetesSVC, instance.KubernetesNS), port, instance.RootPassword), "GET", nil, nil)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"success": false, "message": "请求agent失败"})
		return
	}
	var jsonData map[string]interface{}
	err = json.Unmarshal([]byte(dt), &jsonData)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"success": false, "message": "反序列化agent数据失败"})
		return
	}

	if jsonData["success"].(bool) == false {
		c.AbortWithStatusJSON(200, gin.H{"success": false, "message": fmt.Sprintf("数据库无法链接，请排查用户名，密码及网络:%s", jsonData["message"].(string))})
		return
	}

	//以上校验都通过，开始创建实例
	database.DB.Create(&instance)

	// 创建secret
	secret := client.GetSecret("dms-mysql-coldbak", kubernetesInfo.Namespace)
	if secret == nil {
		dataSecret := util.CreateSecretData()
		secret = client.CreateSecret("dms-mysql-coldbak", kubernetesInfo.Namespace, dataSecret)
		if secret == nil {
			c.AbortWithStatusJSON(200, gin.H{"success": false, "message": "创建对应的secret失败"})
			return
		}
	}
	password, _ := b64.StdEncoding.DecodeString(instance.RootPassword)
	secret.Data[fmt.Sprintf("mysql_password_%s", instance.Name)] = password
	if client.UpdateSecret(kubernetesInfo.Namespace, secret) == nil {
		c.AbortWithStatusJSON(200, gin.H{"success": false, "message": "更新对应的secret失败"})
		return
	}

	var coldbakCron model.MysqlColdBakCron
	if instance.MasterID == 0 {

		var limit *int32
		var l int32 = 1
		limit = &l
		var ttl *int32
		var t int32 = 30
		ttl = &t

		// 创建cronjob 随机时间
		rand.Seed(time.Now().Unix())
		hour := 19 + rand.Intn(5)
		minute := rand.Intn(59)

		coldbakCron.Hour = hour
		coldbakCron.Minute = minute

		spec := v1beta1.CronJobSpec{
			SuccessfulJobsHistoryLimit: limit,
			FailedJobsHistoryLimit:     limit,
			ConcurrencyPolicy:          v1beta1.ForbidConcurrent,
			Schedule:                   fmt.Sprintf("%d %d * * *", minute, hour),
			JobTemplate: v1beta1.JobTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{Name: instance.Name, Namespace: kubernetesInfo.Namespace},
				Spec: batchv1.JobSpec{
					TTLSecondsAfterFinished: ttl,
					Template:                util.CreatePodTemplate(instance, "CronJob"),
				},
			},
		}
		if client.CreateCronJob(instance.Name, kubernetesInfo.Namespace, spec) == nil {
			c.AbortWithStatusJSON(200, gin.H{"success": false, "message": "创建实例，创建cronjob失败"})
			return
		}
	}
	if instance.MasterID == 0 {
		coldbakCron.MysqlInstanceID = instance.ID
		coldbakCron.CreateAt = time.Now()
		database.DB.Create(&coldbakCron)
	}
	data, _ := json.Marshal(instance)
	//操作日志
	go model.CreateOperate(instance.UserID, "MYSQL", "CREATE_INSTANCE", fmt.Sprintf("创建MYSQL实例，name: [%s], namespace:[%s].", instance.Name, kubernetesInfo.Namespace), string(data))

	c.AbortWithStatusJSON(201, gin.H{"success": true})
	return
}

func MysqlInstanceColdbakCronIds(c *gin.Context) {
	var coldbakCrons []model.MysqlColdBakCron
	database.DB.Select("instance_id").Find(&coldbakCrons)
	ids := []int{}
	for _, cron := range coldbakCrons {
		ids = append(ids, cron.MysqlInstanceID)
	}
	c.AbortWithStatusJSON(200, gin.H{"success": true, "data": ids})
	return
}

func MysqlInstanceDelete(c *gin.Context) {
	instanceID, _ := strconv.Atoi(c.Param("instance_id"))
	var instance model.MysqlInstance
	database.DB.Where("is_delete = ?", false).Preload("KubernetesInfo").First(&instance, instanceID)
	// 删除cronjob
	client := util.KubernetesClient{Host: instance.KubernetesInfo.ClusterURL, Port: instance.KubernetesInfo.Port, Token: instance.KubernetesInfo.Token, Scheme: "https"}
	if client.GetCronJob(instance.Name, instance.KubernetesInfo.Namespace) != nil {
		if !client.DeleteCronJob(instance.Name, instance.KubernetesInfo.Namespace) {
			c.AbortWithStatusJSON(204, gin.H{"success": false, "message": "删除cronjob失败"})
			return
		}
	}
	// if client.GetSecret(instance.Name, instance.KubernetesInfo.Namespace) != nil {
	// 	if !client.DeleteSecret(instance.Name, instance.KubernetesInfo.Namespace) {
	// 		c.AbortWithStatusJSON(200, gin.H{"success": false, "message": "删除secret失败"})
	// 		return
	// 	}
	// }

	database.DB.Model(&instance).Updates(model.MysqlInstance{IsDelete: true})
	data, _ := json.Marshal(instance)
	go model.CreateOperate(sessions.Default(c).Get("user").(model.User).ID, "MYSQL", "DELETE_INSTANCE", fmt.Sprintf("删除YSQL实例，name: [%s], namespace:[%s].", instance.Name, instance.KubernetesInfo.Namespace), string(data))
	c.AbortWithStatusJSON(204, gin.H{"success": true, "message": "未发现相关cronjob或者secret"})
	return
}

func MysqlColdBakList(c *gin.Context) {
	pageSizeDefault, _ := runtime.Cfg.GetValue("utils", "PAGE_SIZE")

	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", pageSizeDefault))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

	clusterID := c.DefaultQuery("cluster_id", "")
	search := c.DefaultQuery("search", "")
	coldbakDate := c.DefaultQuery("date", "")

	var total int
	var coldbaks []model.MysqlColdBak
	countDB := database.DB.Order("id desc").Model(&model.MysqlColdBak{}).Joins("inner join mysql_instance on mysql_instance.id = mysql_coldbak.instance_id").Where("mysql_coldbak.create_date = ?", coldbakDate)
	dataDB := database.DB.Order("id desc").Offset((page-1)*pageSize).Limit(pageSize).Preload("User").Preload("MysqlInstance").Preload("MysqlInstance.KubernetesInfo").Joins("inner join mysql_instance on mysql_instance.id = mysql_coldbak.instance_id").Where("mysql_coldbak.create_date = ?", coldbakDate)

	if clusterID != "" {
		countDB = countDB.Where("mysql_instance.kubernetes_id = ?", clusterID)
		dataDB = dataDB.Where("mysql_instance.kubernetes_id = ?", clusterID)
	}
	if search != "" {
		countDB = countDB.Where("mysql_instance.name like ? or mysql_instance.alias like ?", fmt.Sprintf("%%%s%%", search), fmt.Sprintf("%%%s%%", search))
		dataDB = dataDB.Where("mysql_instance.name like ? or mysql_instance.alias like ?", fmt.Sprintf("%%%s%%", search), fmt.Sprintf("%%%s%%", search))
	}

	countDB.Count(&total)
	dataDB.Find(&coldbaks)

	c.AbortWithStatusJSON(200, gin.H{"success": true, "total": total, "data": coldbaks, "page": page, "page_size": pageSize})
	return
}

func MysqlColdBakRestoreList(c *gin.Context) {
	pageSizeDefault, _ := runtime.Cfg.GetValue("utils", "PAGE_SIZE")

	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", pageSizeDefault))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

	clusterID := c.DefaultQuery("cluster_id", "")
	search := c.DefaultQuery("search", "")

	var total int
	var coldbaklogs []model.MysqlColdBakRestoreLog
	countDB := database.DB.Order("id desc").Model(&model.MysqlColdBakRestoreLog{}).Joins("inner join mysql_instance on mysql_instance.id = mysql_coldbak_restore_log.instance_id")
	dataDB := database.DB.Order("id desc").Offset((page - 1) * pageSize).Limit(pageSize).Preload("User").Preload("MysqlInstance").Preload("MysqlColdBak").Preload("MysqlInstance.KubernetesInfo").Joins("inner join mysql_instance on mysql_instance.id = mysql_coldbak_restore_log.instance_id")

	if clusterID != "" {
		countDB = countDB.Where("mysql_instance.kubernetes_id = ?", clusterID)
		dataDB = dataDB.Where("mysql_instance.kubernetes_id = ?", clusterID)
	}
	if search != "" {
		countDB = countDB.Where("mysql_instance.name like ? or mysql_instance.alias like ?", fmt.Sprintf("%%%s%%", search), fmt.Sprintf("%%%s%%", search))
		dataDB = dataDB.Where("mysql_instance.name like ? or mysql_instance.alias like ?", fmt.Sprintf("%%%s%%", search), fmt.Sprintf("%%%s%%", search))
	}

	countDB.Count(&total)
	dataDB.Find(&coldbaklogs)

	c.AbortWithStatusJSON(200, gin.H{"success": true, "total": total, "data": coldbaklogs, "page": page, "page_size": pageSize})
	return
}

func MysqlColdBakClearCronJob(c *gin.Context) {
	instanceID, _ := strconv.Atoi(c.Param("instance_id"))
	var instance model.MysqlInstance
	database.DB.Where("is_delete = ?", false).Preload("KubernetesInfo").First(&instance, instanceID)

	if instance.MasterID > 0 {
		c.AbortWithStatusJSON(200, gin.H{"success": false, "message": "只能对主库或者单实例进行清理冷备"})
		return
	}

	// 清理冷备job
	client := util.KubernetesClient{Host: instance.KubernetesInfo.ClusterURL, Port: instance.KubernetesInfo.Port, Token: instance.KubernetesInfo.Token, Scheme: "https"}

	if client.DeleteCronJob(instance.Name, instance.KubernetesInfo.Namespace) {
		database.DB.Delete(model.MysqlColdBakCron{}, "instance_id = ?", instanceID)
		c.AbortWithStatusJSON(201, gin.H{"success": true})
		return
	}

	c.AbortWithStatusJSON(200, gin.H{"success": false, "message": "删除CronJob失败"})
	return
}

func MysqlColdBakCreateCronJob(c *gin.Context) {
	instanceID, _ := strconv.Atoi(c.Param("instance_id"))
	var instance model.MysqlInstance
	database.DB.Where("is_delete = ?", false).Preload("KubernetesInfo").First(&instance, instanceID)

	if instance.MasterID > 0 {
		c.AbortWithStatusJSON(200, gin.H{"success": false, "message": "只能对主库或者单实例进行冷备"})
		return
	}

	// 创建冷备job
	client := util.KubernetesClient{Host: instance.KubernetesInfo.ClusterURL, Port: instance.KubernetesInfo.Port, Token: instance.KubernetesInfo.Token, Scheme: "https"}
	// 创建secret
	secret := client.GetSecret("dms-mysql-coldbak", instance.KubernetesInfo.Namespace)
	if secret == nil {
		dataSecret := util.CreateSecretData()
		secret = client.CreateSecret("dms-mysql-coldbak", instance.KubernetesInfo.Namespace, dataSecret)
		if secret == nil {
			c.AbortWithStatusJSON(200, gin.H{"success": false, "message": "创建对应的secret失败"})
			return
		}
	}
	password, _ := b64.StdEncoding.DecodeString(instance.RootPassword)
	secret.Data[fmt.Sprintf("mysql_password_%s", instance.Name)] = password
	if client.UpdateSecret(instance.KubernetesInfo.Namespace, secret) == nil {
		c.AbortWithStatusJSON(200, gin.H{"success": false, "message": "更新对应的secret失败"})
		return
	}

	var limit *int32
	var l int32 = 1
	limit = &l
	var ttl *int32
	var t int32 = 30
	ttl = &t

	// 创建cronjob 随机时间
	rand.Seed(time.Now().Unix())
	hour := 19 + rand.Intn(5)
	minute := rand.Intn(59)

	coldbakCron := model.MysqlColdBakCron{
		Hour:            hour,
		Minute:          minute,
		CreateAt:        time.Now(),
		MysqlInstanceID: instance.ID,
	}

	spec := v1beta1.CronJobSpec{
		SuccessfulJobsHistoryLimit: limit,
		FailedJobsHistoryLimit:     limit,
		ConcurrencyPolicy:          v1beta1.ForbidConcurrent,
		Schedule:                   fmt.Sprintf("%d %d * * *", minute, hour),
		JobTemplate: v1beta1.JobTemplateSpec{
			ObjectMeta: metav1.ObjectMeta{Name: instance.Name, Namespace: instance.KubernetesInfo.Namespace},
			Spec: batchv1.JobSpec{
				TTLSecondsAfterFinished: ttl,
				Template:                util.CreatePodTemplate(instance, "CronJob"),
			},
		},
	}

	if client.CreateCronJob(instance.Name, instance.KubernetesInfo.Namespace, spec) != nil {
		database.DB.Create(&coldbakCron)

		data, _ := json.Marshal(instance)
		go model.CreateOperate(sessions.Default(c).Get("user").(model.User).ID, "MYSQL", "CREATE_COLDBAK", fmt.Sprintf("创建冷备CronJob任务，name: [%s], namespace:[%s].", instance.Name, instance.KubernetesInfo.Namespace), string(data))

		c.AbortWithStatusJSON(201, gin.H{"success": true, "message": "冷备CronJob创建成功"})
		return
	}
	c.AbortWithStatusJSON(200, gin.H{"success": false, "message": "创建冷备CrobJob失败"})
	return
}

func MysqlColdBakCreate(c *gin.Context) {
	instanceID, _ := strconv.Atoi(c.Param("instance_id"))
	var instance model.MysqlInstance
	//// 查询指定的某条记录(仅当主键为整型时可用)
	database.DB.Where("is_delete = ?", false).Preload("KubernetesInfo").First(&instance, instanceID)

	if instance.MasterID > 0 {
		c.AbortWithStatusJSON(200, gin.H{"success": false, "message": "只能对主库或者单实例进行冷备"})
		return
	}

	// 创建冷备job
	client := util.KubernetesClient{Host: instance.KubernetesInfo.ClusterURL, Port: instance.KubernetesInfo.Port, Token: instance.KubernetesInfo.Token, Scheme: "https"}

	if client.GetJob(instance.Name, instance.KubernetesInfo.Namespace) != nil {
		c.AbortWithStatusJSON(200, gin.H{"success": false, "message": "当前存在正在进行中的Job，请稍后再试"})
		return
	}

	// 创建secret
	secret := client.GetSecret("dms-mysql-coldbak", instance.KubernetesInfo.Namespace)
	if secret == nil {
		dataSecret := util.CreateSecretData()
		secret = client.CreateSecret("dms-mysql-coldbak", instance.KubernetesInfo.Namespace, dataSecret)
		if secret == nil {
			c.AbortWithStatusJSON(200, gin.H{"success": false, "message": "创建对应的secret失败"})
			return
		}
	}
	password, _ := b64.StdEncoding.DecodeString(instance.RootPassword)
	secret.Data[fmt.Sprintf("mysql_password_%s", instance.Name)] = password
	if client.UpdateSecret(instance.KubernetesInfo.Namespace, secret) == nil {
		c.AbortWithStatusJSON(200, gin.H{"success": false, "message": "更新对应的secret失败"})
		return
	}

	var ttl *int32
	var l int32 = 30
	ttl = &l
	spec := batchv1.JobSpec{
		TTLSecondsAfterFinished: ttl,
		Template:                util.CreatePodTemplate(instance, "Job"),
	}

	if client.CreateJob(instance.Name, instance.KubernetesInfo.Namespace, spec) != nil {
		data, _ := json.Marshal(instance)
		go model.CreateOperate(sessions.Default(c).Get("user").(model.User).ID, "MYSQL", "CREATE_COLDBAK", fmt.Sprintf("创建冷备任务，name: [%s], namespace:[%s].", instance.Name, instance.KubernetesInfo.Namespace), string(data))

		c.AbortWithStatusJSON(201, gin.H{"success": true, "message": "冷备Job创建成功，请等待冷备结果"})
		return
	}
	c.AbortWithStatusJSON(200, gin.H{"success": false, "message": "创建冷备Job失败"})
	return
}

func MysqlColdBakCallback(c *gin.Context) {
	postDataByte, _ := ioutil.ReadAll(c.Request.Body)
	var postData map[string]interface{}
	err := json.Unmarshal(postDataByte, &postData)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"success": false, "message": err.Error()})
		return
	}

	filePath := postData["file_path"].(string)
	timeConsumption := int(postData["time_consumption"].(float64))
	trigger := postData["trigger"].(string)
	size := postData["size"].(string)
	message := postData["message"].(string)
	logPath := postData["log_path"].(string)
	resultFlag := postData["flag"].(string)
	instanceID, _ := strconv.Atoi(postData["id"].(string))
	userID, _ := strconv.Atoi(postData["user_id"].(string))
	coldbakUUID := postData["coldbak_uuid"].(string)

	coldbak := model.MysqlColdBak{
		Trigger:         trigger,
		FilePath:        filePath,
		LogPath:         logPath,
		Size:            size,
		TimeConsumption: timeConsumption,
		Message:         message,
		ResultFlag:      resultFlag,
		UserID:          userID,
		MysqlInstanceID: instanceID,
		CreateAt:        time.Now(),
		CreateDate:      util.MysqlTime(),
	}

	var lastColdbak model.MysqlColdBak
	database.DB.Where("instance_id = ?", instanceID).Order("id desc").First(&lastColdbak)
	database.DB.Create(&coldbak)

	var instance model.MysqlInstance
	database.DB.Preload("KubernetesInfo").First(&instance, instanceID)

	if util.FileSize(lastColdbak.Size) > util.FileSize(size) {
		// 备份文件变小
		util.SendMail([]string{"leon.chen@cloudminds.com", "shilei.wang@cloudminds.com"}, "冷备备份文件大小异常", fmt.Sprintf("集群[%s],冷备实例[%s]本次冷备结果存在大小异常，本次冷备文件大小[%s], 前一次冷备文件大小[%s]，请确认。", instance.KubernetesInfo.Alias, instance.Name, size, lastColdbak.Size))
	}

	if coldbak.Trigger == "Job" {
		client := util.KubernetesClient{Host: instance.KubernetesInfo.ClusterURL, Port: instance.KubernetesInfo.Port, Token: instance.KubernetesInfo.Token, Scheme: "https"}
		job := client.GetJob(fmt.Sprintf("%s", instance.Name), instance.KubernetesInfo.Namespace)
		if job != nil {
			client.DeleteJob(fmt.Sprintf("%s", instance.Name), instance.KubernetesInfo.Namespace)
		}
	}

	data, _ := json.Marshal(coldbak)
	msg := ""
	if coldbak.Trigger == "Job" {
		msg = fmt.Sprintf("冷备任务回调，uuid: [%s]，name: [%s], namespace:[%s].", coldbakUUID, instance.Name, instance.KubernetesInfo.Namespace)
	} else {
		msg = fmt.Sprintf("冷备任务回调，name: [%s], namespace:[%s].", instance.Name, instance.KubernetesInfo.Namespace)
	}
	go model.CreateOperate(userID, "MYSQL", "COLDBAK_CALLBACK", msg, string(data))

	c.AbortWithStatusJSON(200, gin.H{"success": true})
	return
}

func MysqlColdBakRestore(c *gin.Context) {
	coldbakID, _ := strconv.Atoi(c.Param("coldbak_id"))
	var coldbak model.MysqlColdBak
	database.DB.Preload("MysqlInstance").Preload("MysqlInstance.KubernetesInfo").First(&coldbak, coldbakID)

	var restore model.MysqlColdBakRestoreLog
	database.DB.Where("result_flag = ? and instance_id = ?", "DOING", coldbak.MysqlInstanceID).First(&restore)
	if restore.ID > 0 {
		c.AbortWithStatusJSON(201, gin.H{"success": false, "message": "存在未执行完成的冷备恢复任务，请等待完成后再操作"})
		return
	}

	// 创建冷备恢复job
	client := util.KubernetesClient{Host: coldbak.MysqlInstance.KubernetesInfo.ClusterURL, Port: coldbak.MysqlInstance.KubernetesInfo.Port, Token: coldbak.MysqlInstance.KubernetesInfo.Token, Scheme: "https"}

	// s3
	password, _ := b64.StdEncoding.DecodeString(coldbak.MysqlInstance.RootPassword)
	s3FilePath := coldbak.FilePath[strings.Index(coldbak.FilePath, "dms/"):]
	localFilePath := fmt.Sprintf("/usr/src/app/files/%s", coldbak.FilePath[strings.LastIndex(coldbak.FilePath, "/")+1:])
	cmdS3 := []string{"/usr/src/app/bin/mc", "-C", "/usr/src/app/.mc", "cp", fmt.Sprintf("dms/%s", s3FilePath), "/usr/src/app/files/"}
	unzipDumps := []string{"gzip", "-df", localFilePath}
	cmdResetMaster := []string{"mysql", "-h", fmt.Sprintf("%s.%s", coldbak.MysqlInstance.KubernetesSVC, coldbak.MysqlInstance.KubernetesNS), fmt.Sprintf("-P%d", coldbak.MysqlInstance.Port), "-uroot", fmt.Sprintf("-p%s", password), "-e", "reset master;"}
	cmdRestore := []string{"mysql", "-h", fmt.Sprintf("%s.%s", coldbak.MysqlInstance.KubernetesSVC, coldbak.MysqlInstance.KubernetesNS), fmt.Sprintf("-P%d", coldbak.MysqlInstance.Port), "-uroot", fmt.Sprintf("-p%s", password), "-e", fmt.Sprintf("source %s;", strings.ReplaceAll(localFilePath, ".gz", ""))}
	cmdClear := []string{"rm", "-rf", strings.ReplaceAll(localFilePath, ".gz", "")}

	coldbakRestoreLog := model.MysqlColdBakRestoreLog{
		TimeConsumption: 0,
		Message:         "",
		ResultFlag:      "DOING",
		UserID:          sessions.Default(c).Get("user").(model.User).ID,
		MysqlColdBakID:  coldbak.ID,
		MysqlInstanceID: coldbak.MysqlInstance.ID,
		CreateAt:        time.Now(),
	}

	data, _ := json.Marshal(coldbakRestoreLog)
	go model.CreateOperate(sessions.Default(c).Get("user").(model.User).ID, "MYSQL", "RESTORE_COLDBAK", fmt.Sprintf("创建冷备恢复任务，name: [%s], namespace:[%s].", coldbak.MysqlInstance.Name, coldbak.MysqlInstance.KubernetesInfo.Namespace), string(data))
	database.DB.Create(&coldbakRestoreLog)

	// 冷备恢复job
	agentPod, _ := util.GetAgent(coldbak.MysqlInstance.KubernetesInfo.Namespace, client)
	go job.ColdbakRestoring(agentPod, coldbak.MysqlInstance.KubernetesInfo.Namespace, [][]string{cmdS3, unzipDumps, cmdResetMaster, cmdRestore, cmdClear}, &coldbakRestoreLog, &client)

	c.AbortWithStatusJSON(201, gin.H{"success": true})
	return
}

func MysqlInstanceConnection(c *gin.Context) {
	instanceID, _ := strconv.Atoi(c.Param("instance_id"))
	var instance model.MysqlInstance
	database.DB.Preload("KubernetesInfo").First(&instance, instanceID)

	client := util.KubernetesClient{Host: instance.KubernetesInfo.ClusterURL, Port: instance.KubernetesInfo.Port, Token: instance.KubernetesInfo.Token, Scheme: "https"}

	_, agentPort := util.GetAgent(instance.KubernetesInfo.Namespace, client)
	data, err := util.HttpRequest(fmt.Sprintf("http://%s:%d/api/mysql/connect/check?host=%s&port=%d&username=root&password=%s", instance.KubernetesInfo.ClusterURL, agentPort, fmt.Sprintf("%s.%s", instance.KubernetesSVC, instance.KubernetesNS), instance.Port, instance.RootPassword), "GET", nil, nil)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"success": false, "message": "请求agent失败"})
		return
	}
	var jsonData map[string]interface{}
	err = json.Unmarshal([]byte(data), &jsonData)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"success": false, "message": "反序列化agent数据失败"})
		return
	}
	if jsonData["success"].(bool) {
		c.AbortWithStatusJSON(200, gin.H{"success": true, "data": jsonData["success"].(bool)})
	} else {
		c.AbortWithStatusJSON(200, gin.H{"success": true, "data": jsonData["success"].(bool), "message": jsonData["message"].(string)})
	}
	return
}

func MysqlInstanceEdit(c *gin.Context) {
	instanceID, _ := strconv.Atoi(c.Param("instance_id"))

	postDataByte, _ := ioutil.ReadAll(c.Request.Body)
	var postData map[string]interface{}
	err := json.Unmarshal(postDataByte, &postData)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"success": false, "message": err.Error()})
		return
	}

	alias := postData["alias"].(string)
	port := int(postData["port"].(float64))
	rootPassword := postData["root_password"].(string)

	var instance model.MysqlInstance
	database.DB.Preload("KubernetesInfo").First(&instance, instanceID)

	client := util.KubernetesClient{Host: instance.KubernetesInfo.ClusterURL, Port: instance.KubernetesInfo.Port, Token: instance.KubernetesInfo.Token, Scheme: "https"}

	_, agentPort := util.GetAgent(instance.KubernetesInfo.Namespace, client)
	data, err := util.HttpRequest(fmt.Sprintf("http://%s:%d/api/mysql/connect/check?host=%s&port=%d&username=root&password=%s", instance.KubernetesInfo.ClusterURL, agentPort, fmt.Sprintf("%s.%s", instance.KubernetesSVC, instance.KubernetesNS), port, b64.StdEncoding.EncodeToString([]byte(rootPassword))), "GET", nil, nil)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"success": false, "message": "请求agent失败"})
		return
	}
	var jsonData map[string]interface{}
	err = json.Unmarshal([]byte(data), &jsonData)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"success": false, "message": "反序列化agent数据失败"})
		return
	}

	if jsonData["success"].(bool) == false {
		c.AbortWithStatusJSON(200, gin.H{"success": false, "message": fmt.Sprintf("数据库无法链接，请排查用户名，密码及网络:%s", jsonData["message"].(string))})
		return
	}

	// 更新secret
	secret := client.GetSecret("dms-mysql-coldbak", instance.KubernetesInfo.Namespace)
	secret.Data[fmt.Sprintf("mysql_password_%s", instance.Name)] = []byte(rootPassword)
	if client.UpdateSecret(instance.KubernetesInfo.Namespace, secret) == nil {
		c.AbortWithStatusJSON(200, gin.H{"success": false, "message": "更新对应的secret失败"})
		return
	}

	instance.Alias = alias
	instance.Port = int(port)
	instance.RootPassword = b64.StdEncoding.EncodeToString([]byte(rootPassword))
	database.DB.Save(&instance)

	logData, _ := json.Marshal(instance)
	go model.CreateOperate(sessions.Default(c).Get("user").(model.User).ID, "MYSQL", "UPDATE_INSTANCE", fmt.Sprintf("更新Mysql实例，name: [%s], namespace:[%s].", instance.Name, instance.KubernetesInfo.Namespace), string(logData))

	c.AbortWithStatusJSON(200, gin.H{"success": true})
	return
}

// 主从切换
func MysqlInstanceListTree(c *gin.Context) {
	//pageSizeDefault, _ := runtime.Cfg.GetValue("utils", "PAGE_SIZE")
	//
	//search := c.DefaultQuery("search", "")
	//pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", pageSizeDefault))
	//page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	//
	//clusterID := c.DefaultQuery("cluster_id", "")
	////fmt.Printf("MysqlInstanceList clusterID:%v\n", clusterID)
	//
	//var total int
	//var instances []model.MysqlInstance
	//countDB := database.DB.Order("id desc").Model(&model.MysqlInstance{}).Where("is_delete = ?", false)
	//dataDB := database.DB.Order("id desc").Offset((page-1)*pageSize).Limit(pageSize).Preload("User").Preload("KubernetesInfo").Where("is_delete = ?", false)
	////fmt.Printf("MysqlInstanceList countDB:%v\n", countDB)
	////fmt.Printf("MysqlInstanceList dataDB:%v\n", dataDB)
	//
	//if clusterID != "" {
	//	countDB = countDB.Where("kubernetes_id = ?", clusterID)
	//	dataDB = dataDB.Where("kubernetes_id = ?", clusterID)
	//}
	//if search != "" {
	//	countDB = countDB.Where("name like ? or alias like ?", fmt.Sprintf("%%%s%%", search), fmt.Sprintf("%%%s%%", search))
	//	dataDB = dataDB.Where("name like ? or alias like ?", fmt.Sprintf("%%%s%%", search), fmt.Sprintf("%%%s%%", search))
	//}
	//
	//countDB.Count(&total)
	//dataDB.Find(&instances)
	//
	//for _, v := range instances {
	//	fmt.Printf("\nMysqlInstanceList 获取实例列表:\n%+v\n", v)
	//}

	type Basic struct {
		Id       int     `json:"id"`
		Name     string  `json:"name"`
		Children []Basic `json:"children,omitempty"` //值为空时直接忽略,类似递归
	}
	jsonData := []Basic{{
		Id:   1,
		Name: "root",
		Children: []Basic{
			{Id: 2, Name: "master1", Children: []Basic{{Id: 3, Name: "slave1"}, {Id: 4, Name: "slave2"}}},
			{Id: 5, Name: "master2", Children: []Basic{{Id: 6, Name: "slave1"}, {Id: 7, Name: "slave2"}}},
		},
	}}

	str, err := json.Marshal(jsonData)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(str))

	//获取MySQL实例列表

	c.AbortWithStatusJSON(200, gin.H{"success": true, "data": jsonData})
	return
}
