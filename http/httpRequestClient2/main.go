package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var mytoken string

//封装Http请求
func HttpRequest(apiURL string, method string, headers map[string]string, data interface{}) (string, error) {
	client := &http.Client{
		Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, DisableKeepAlives: true},
		//Timeout: 30,
	}

	requestData, _ := json.Marshal(data)
	req, err := http.NewRequest(method, apiURL, bytes.NewBuffer(requestData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "JWT "+mytoken)
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	if err != nil {
		return "", err
	}

	response, err := client.Do(req)
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()
	if response.StatusCode == 200 || response.StatusCode == 201 {
		return string(body), nil
	} else {
		return "", fmt.Errorf("%s", string(body))
	}
}

func main() {
	//获取大禹认证token
	//cmdbTokenApi := "https://dayu.cloudminds.com/api/v1/cmdb/rest-framework-jwt/token"
	cmdbTokenApi := "https://dayu-dev.cloudminds.com/api/v1/cmdb/rest-framework-jwt/token"
	account := map[string]string{
		"username": "xiaobing.song",
		"password": "SXB@cloud",
	}
	jwtData, _ := HttpRequest(cmdbTokenApi, "POST", nil, account)
	fmt.Printf("获取Token返回结果:\n%+v\n", jwtData)

	jwtDataMap := map[string]string{}
	json.Unmarshal([]byte(jwtData), &jwtDataMap)
	mytoken = jwtDataMap["token"]
	fmt.Printf("Token:\n%+v\n", jwtDataMap["token"])

	postData := map[string]interface{}{"file_path": "https://s3.harix.iamidata.com/dms/cl-mng-cloudpepper-v3-dit-cl-mng-mysql/20201014063403/cl-mng-cloudpepper-v3-dit-cl-msql.binlog-on.dump.gz", "time_consumption": 10.485089, "size": "967.030 KB", "coldbak_uuid": "20201014063324", "flag": "DONE", "trigger": "Job", "message": "[2020-10-14 06:34:03]INFO: \\u672c\\u6b21\\u51b7\\u5907\\u53c2\\u6570\\uff1auser: root, port: 3306, instance_name: cl-mng-cloudpepper-v3-dit-cl-mng-mysql, log_push_url: https://dayu.cloudminds.com/api/v1/dms/mysql/coldbak/k\n[2020-10-14 06:34:03]INFO\\uff1a/logs/cl-mng-cloudpepper-v3-dit-cl-mng-mysql/20201014063403.log\n[2020-10-14 06:34:03]INFO: bin log status: (True, 'ON')\n[2020-10-14 06:34:03INFO: mysqldump -hcl-mng-mysql.cl-mng-cloudpepper-v3-dit -P3306 --all-databases --compress --flush-logs --flush-privileges --master-data=2 --routines --triggers --events --single-transaction --dump-date --log-error=/MYSQL/mysql_dump.log.err |gzip > /MYSQL/cl-mng-cloudpepper-v3-dit-cl-mng-mysql/20201014063403/cl-mng-cloudpepper-v3-dit-cl-mng-mysql.binlog-on.dump.gz\n[2020-10-14 06:34:03]INFO: \\u51b7\\u5907\\u6587\\u4ef6\\u8def\\u5f84: /MYSQL/cl-mng-cloudpepper-v3-dit-cl-mng-mysql/20201014063403/cl-mng-cloudpepper-v3-dit-cl-mng-mysql.binlog-p.gz\n[2020-10-14 06:34:13]INFO: \\u6267\\u884c\\u547d\\u4ee4\\u7ed3\\u679c: CompletedProcess(args='mysqldump -hcl-mng-mysql.cl-mng-cloudpepper-v3-dit -P3306 -uroot -p123456 --all-databases --compress --logs --flush-privileges --master-data=2 --routines --triggers --events --single-transaction --dump-date --log-error=/MYSQL/mysql_dump.log.err |gzip > /MYSQL/cl-mng-cloudpepper-v3-dit-cl-mng-mysql/20201014063403/cl-mng-cloudpepper-v3-dit-cl-mng-mysql.binlog-on.dump.gz', returncode=0, stdout=b'', stderr=b'Warning: Using a password on the command line interface can be insecure.\n')\n[2020-10-14 06:34:13]INFO: \\u8017\\u65f6: 10.485089\n[2020-10-14 06:34:13]INFO: \\u51b7\\u5907\\u6587\\u4ef6\\u5927\\u5c0f: 967.030 KB\n[2020-10-14 06:34:13]INFO: \\u51b7\\u5907\\u7ed3\\u679c[2020-10-14 06:34:13]INFO: \\u51b7\\u5907\\u547d\\u4ee4\\u6267\\u884c\\u9519\\u8bef\\u65e5\\u5fd7: \n[2020-10-14 06:34:13]INFO: \\u4e0a\\u4f20\\u5907\\u4efd\\u6587\\u4ef6\\u5230Ceph S3: https://s3.harix.iamidata.com/dms/cl-mng-cloudpepper-v3-dit-cl-mng-mysqcl-mng-cloudpepper-v3-dit-cl-mng-mysql.binlog-on.dump.gz\n[2020-10-14 06:34:14]INFO: data\\u6587\\u4ef6\\u4e0a\\u4f20\\u6267\\u884c\\u7ed3\\u679c: CompletedProcess(args='/usr/bin/mc --config-dir /root/.mc cp /MYSg-cloudpepper-v3-dit-cl-mng-mysql/20201014063403/cl-mng-cloudpepper-v3-dit-cl-mng-mysql.binlog-on.dump.gz dms/dms/cl-mng-cloudpepper-v3-dit-cl-mng-mysql/20201014063403/cl-mng-cloudpepper-v3-dit-cl-mng-mysql.binlog-on.dump.gz', returncode=0, stdout=b'`/MYSQL/cl-mng-cloudpepper-v3-dit-cl-mng-mysql/20201014063403/cl-mng-cloudpepper-v3-dit-cl-mng-mysql.binlog-on.dump.gz` -> `dms/dms/cl-mng-cloudpepper-v3-dit-cl-mng-mysql/20201014063403/cl-mng-cloudpepper-v3-dit-cl-mng-mysql.binlog-on.dump.gz`\nTotal: 967.03 KiB, Transferred: 967.03 KiB, Speed: 21.15 MiB/s\n', stderr=b'')\n[2020-10-14 06:34:14]INFO: \\u4e0a\\u4f20\\u65e5\\u5fd7\\u6587\\u4ef6\\u5230Ceph S3: https://s3.harix.iamidata.com/dms/cl-mng-cloudpepper-v3-dit-cysql/20201014063403/cl-mng-cloudpepper-v3-dit-cl-mng-mysql.log\n[2020-10-14 06:34:14]INFO: log\\u6587\\u4ef6\\u4e0a\\u4f20\\u6267\\u884c\\u7ed3\\u679c: CompletedProcess(args='/usr/bin/mc --config-dir /root/.mc cpl-mng-cloudpepper-v3-dit-cl-mng-mysql/20201014063403.log dms/dms/cl-mng-cloudpepper-v3-dit-cl-mng-mysql/20201014063403/cl-mng-cloudpepper-v3-dit-cl-mng-mysql.log', returncode=0, stdout=b'`/logs/cl-mng-cloudpepper-v3-dit-cl-mng-mysql/20201014063403.log` -> `dms/dms/cl-mng-cloudpepper-v3-dit-cl-mng-mysql/20201014063403/cl-mng-cloudpepper-v3-dit-cl-mng-mysql.log`\nTotal: 2.71 KiB, Transferred: 2.71 KiB, Speed: 321.52 KiB/s\n', stderr=b'')\n", "id": "47", "user_id": "1", "log_path": "/logs/cl-mng-cloudpepper-v3-dit-cl-mng-mysql/20201014063403.log"}

	getData, _ := HttpRequest("https://dayu-dev.cloudminds.com/api/v1/dms/mysql/coldbak/callback", "POST", nil, postData)
	fmt.Printf("POST数据:\n%+v\n", getData)

}
