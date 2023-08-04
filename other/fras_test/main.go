package main

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Employee struct {
	Id            int    `orm:"column(id);auto" description:"id号码"`
	Live          string `orm:"column(live)" description:"是否激活状态，删除的时候标记为No"`
	Eid           string `orm:"column(eid);size(20);null" description:"工号Employee ID"`
	IDNo          string `orm:"column(idno);size(20);null" description:"身份证号码"`
	Name          string `orm:"column(name);size(20);null" description:"中文名称"`
	DepartmentID  int    `orm:"column(departmentID);null" description:"部门ID，和部门库内联"`
	Gender        string `orm:"column(gender)"`
	Birthday      string `orm:"column(birthday);null" description:"出生日期"`
	Phone         string `orm:"column(phone);size(18);null" description:"电话号码"`
	PhotoPath     string `orm:"column(photoPath);size(255);null" description:"照片保存地址"`
	ButiPhotoPath string `orm:"column(butiPhotoPath);size(255)" description:"美颜照片保存地址"`
	Position      string `orm:"column(position)" description:"位置"`
	// Facesetid     string    `orm:"column(facesetid)" description:"人脸库id"`
	Pubtime time.Time `orm:"column(pubtime);type(timestamp);auto_now" description:"数据更新时间"`
	// GroupId       string    `orm:"column(groupid);size(200)" description:"所属组id"`
	Tenantcode string `orm:"column(tenantcode);size(200)" description:"租户"`
	S3url1     string `json:"s3url1" form:"s3url1" description:"人脸图片存储路径"`
	S3url2     string `json:"s3url2" form:"s3url2" description:"人脸展示图片存储路径"`
	WelcomeMsg string `orm:"column(welcomemsg)" description:"欢迎语"`
}

func main() {
	var (
		err      error
		employee []Employee
		//vv       Employee
	)

	//orm.Debug = true

	//var mu sync.Mutex

	begintime := time.Now().Unix()
	fmt.Println("fras test---------------------------->start：", begintime)
	if err = orm.RegisterDriver("mysql", orm.DRMySQL); err != nil {
		fmt.Println("RegisterDriver error：", err)
	} else {
		//if err = orm.RegisterDataBase("default", "mysql", "sunny:Dt1210k#@tcp(172.16.31.96:30677)/smartomp_cms?charset=utf8mb4&loc=Local"); err != nil {   //生产环境231
		//if err = orm.RegisterDataBase("default", "mysql", "sunny:Dt1210k#@tcp(172.16.13.135:30677)/smartomp_cms?charset=utf8mb4&loc=Local"); err != nil {



		if err = orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/smartomp_cms?charset=utf8mb4&loc=Local"); err != nil {
			fmt.Println("RegisterDataBase error：", err)
		} else {
			//db, err := orm.GetDB("default")
			//db.SetMaxIdleConns(0)
			////db.SetConnMaxLifetime(-1)
			//db.SetConnMaxLifetime(3*time.Second)

			o := orm.NewOrm()

			//orm.
			num, err := o.Raw("select * from employee").QueryRows(&employee)
			//num, err := o.Raw("select * from employee limit 50").QueryRows(&employee)

			if err != nil {
				fmt.Println("raw error：", err)
			} else {
				fmt.Println("employee total:", num)
				taskLimit := int(num)
				//taskLimit := 500
				limiter := make(chan struct{}, taskLimit)
				res := make(chan []string, num)
				defer close(limiter)
				count := 1
				logs.Info("员工数量:", len(employee))
				for _, v := range employee {
					logs.Info("计数器count=", count)
					count += 1
					limiter <- struct{}{}  //写入空对象
					go func(v Employee) {
						defer func() {
							<-limiter  //读一个空对象
							//obj := <-limiter  //读一个空对象
							//logs.Info("读出对象:", obj, count)
						}()
						var vv Employee
						//mu.Lock()
						err := o.Raw("SELECT * FROM employee WHERE id = ?", v.Id).QueryRow(&vv)
						if err == nil {
							logs.Info("read success：", v.Id)
						} else {
							logs.Info("read error", v.Id, err)
							println(err)
						}
						//mu.Unlock()
					}(v)
				}
				//等待所有goroutine完成
				for i := 0; i < taskLimit; i++ {
					limiter <- struct{}{}
				}
				close(res)
			}
		}
	}
	endtime := time.Now().Unix()
	logs.Info("fras test---------------------------->end：", begintime, endtime, endtime-begintime)
}
