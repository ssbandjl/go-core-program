package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//封装Http请求
func HttpRequest(apiURL string, method string, headers map[string]string, data interface{}) (string, error) {
	client := &http.Client{
		Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, DisableKeepAlives: true},
		//Timeout: 30,
	}

	requestData, _ := json.Marshal(data)
	req, err := http.NewRequest(method, apiURL, bytes.NewBuffer(requestData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "JWT "+"eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyX2lkIjoyLCJ1c2VybmFtZSI6InhpYW9iaW5nLnNvbmciLCJleHAiOjE2MDE0NzkyNzEsImVtYWlsIjoieGlhb2Jpbmcuc29uZ0BjbG91ZG1pbmRzLmNvbSIsIm9yaWdfaWF0IjoxNjAxNDM2MDcxfQ.po4xU0IjmlRns7L__NgLPlDotgq3CFMcuYcPPQp8LeU")
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
	cmdbTokenApi := "http://localhost:8000/api/v1/cmdb/rest-framework-jwt/token"
	account := map[string]string{
		"username": "xiaobing.song",
		"password": "SXB@cloud",
	}
	jwtData, _ := HttpRequest(cmdbTokenApi, "POST", nil, account)
	fmt.Printf("获取Token返回结果:\n%+v\n", jwtData)

	jwtDataMap := map[string]string{}
	json.Unmarshal([]byte(jwtData), &jwtDataMap)
	fmt.Printf("Token:\n%+v\n", jwtDataMap["token"])

	getData, _ := HttpRequest("http://localhost:8000/api/v1/cmdb/asset/asset_dashboard", "GET", nil, account)
	fmt.Printf("获取数据:\n%+v\n", getData)

}
