package test_test

import (
	"encoding/json"
	. "github.com/smartystreets/goconvey/convey" //https://github.com/smartystreets/goconvey GoConvey是Golang的测试框架,集成go test, 支持终端和浏览器模式.
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"testing"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Response struct {
	Token string `json:"token"`
}

func createNewsUser(username, password string) *User {
	return &User{username, password}
}

func TestLogin(t *testing.T) {
	Convey("Should be able to login", t, func() {
		user := createNewsUser("jonas", "1234")
		jsondata, _ := json.Marshal(user)
		userData := strings.NewReader(string(jsondata))
		log.Printf("userData:%+v", userData)
		// 这里模拟用户登录, 实际上后台没有使用用户名和密码, 该接口直接返回内部生成的Token
		req, _ := http.NewRequest("GET", "http://localhost:8080/api/", userData)
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		res, _ := client.Do(req)
		//log.Printf("res:%+v", res)
		So(res.StatusCode, ShouldEqual, 200) //对响应码进行断言, 期望得到状态码为200

		Convey("Should be able to parse body", func() { //解析响应体
			body, err := ioutil.ReadAll(res.Body)
			defer res.Body.Close()
			So(err, ShouldBeNil)
			Convey("Should be able to get json back", func() {
				responseData := new(Response)
				err := json.Unmarshal(body, responseData)
				So(err, ShouldBeNil)
				log.Printf("responseData:%s", responseData)
				Convey("Should be able to be authorized", func() {
					token := responseData.Token //提取Token
					log.Printf("token:%s", token)
					// 构造带Token的请求
					req, _ := http.NewRequest("GET", "http://localhost:8080/api/private", nil)
					req.Header.Set("Authorization", "Bearer "+token) //设置认证头
					client = &http.Client{}
					res, _ := client.Do(req)
					body, err := ioutil.ReadAll(res.Body)
					if err != nil {
						log.Printf("Read body failed, %s", err.Error())
					}
					log.Printf("Body:%s", string(body))
					So(res.StatusCode, ShouldEqual, 200)
				})
			})
		})
	})

	//Convey("Should not be able to login with false credentials", t, func() {
	//	user := createNewsUser("jnwfkjnkfneknvjwenv", "wenknfkwnfknfknkfjnwkfenw")
	//	jsondata, _ := json.Marshal(user)
	//	post_data := strings.NewReader(string(jsondata))
	//	req, _ := http.NewRequest("POST", "http://localhost:3000/api/login", post_data)
	//	req.Header.Set("Content-Type", "application/json")
	//	client := &http.Client{}
	//	res, _ := client.Do(req)
	//	So(res.StatusCode, ShouldEqual, 401)
	//})
	//
	//Convey("Should not be able to authorize with false credentials", t, func() {
	//	token := ""
	//	req, _ := http.NewRequest("GET", "http://localhost:3000/api/auth/testAuth", nil)
	//	req.Header.Set("Authorization", "Bearer "+token)
	//	client := &http.Client{}
	//	res, _ := client.Do(req)
	//	So(res.StatusCode, ShouldEqual, 401)
	//})
}
