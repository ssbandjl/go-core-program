package getserviceid

import (
	"encoding/json"
	"fmt"
	_ "io/ioutil"
	"net/http"
	// "io/ioutil"
)

// type CatlogNodes struct {
// 	CatlogNodes []Catlog `json:"catlog_nodes"`
// }

// type Catlog struct {
// 	Product_id string `json: "product_id"`
// 	Quantity   int    `json: "quantity"`
// }

type ServiceDetail struct {
	Kind              string
	Service           string
	ID                string
	Tags              []string
	Meta              map[string]string
	Port              int64
	Address           string
	Weights           map[string]string
	EnableTagOverride string
	CreateIndex       string
	ModifyIndex       string
	ProxyDestination  string
	Connect           string
}

func getserviceid(string service, string consulApiPrefix) {

	// file, _ := ioutil.ReadFile("consul.json")
	// fmt.Printf("\n文件内容:\n%v", string([]byte(file))) //文件(二进制) -> 字节码 -> 字符串
	string getServiceApi
	getServiceApi = consulApiPrefix + "v1/agent/services"

	if service == "" {
		fmt.Printf("服务名不能为空, 使用案例: ./xxx -s auth3-server -a(非必选) http://127.0.0.1:8500/\n")
		return
	}

	resp, err := http.Get(getServiceApi)

	//curl -X GET "https://httpbin.org/get" -H "accept: application/json"

	// fmt.Printf("返回值: %v, 错误信息: %v", resp, err)
	// fmt.Printf("返回值: %+v, 错误信息: %v", resp, err)
	// fmt.Printf("返回值Body: %+v, 错误信息: %v", resp.Body, err)

	// m := make(map[string]interface{})
	m := make(map[string]ServiceDetail)
	// data := ConsulJson2{}

	// _ = json.Unmarshal([]byte(file), &m)
	// file, _ := ioutil.ReadFile("consul.json")
	// _ = json.Unmarshal([]byte(file), &m)

	//将Response.Body中的Json发序列化为结构体
	if err := json.NewDecoder(resp.Body).Decode(&m); err != nil {
		// if err := json.NewDecoder(resp.Body).Decode(&m); err != nil {
		fmt.Println("解析出错: %+v", err)
	}

	// //调式代码，将Post的json数据r.Body绑定的结构体promMessage进行序列化后打印
	data, err := json.Marshal(&m)
	// data, err := json.Marshal(&m)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
		fmt.Printf("\n\n返回的Json数据:\n%v\n", string(data))
	}
	//输出序列化后的结果
	// fmt.Printf("\n\n返回的Json数据:\n%v\n", string(data))

	// for i := 0; i < len(data.CatlogNodes); i++ {
	// 	fmt.Println("Product Id: ", data.CatlogNodes[i].Product_id)
	// 	fmt.Println("Quantity: ", data.CatlogNodes[i].Quantity)
	// }
	// fmt.Printf("\n结构体内容:\n%v", m)
	// fmt.Printf("\n结构体内容:\n%v", m["auth3-server-10.60.21.71-9990"])
	for _, v := range m {
		// fmt.Println(k, v)
		if v.Service == service {
			// if v.Service == "auth3-server" {
			fmt.Printf("\n服务ID:\n%v\n", v.ID)
		}
	}

	fmt.Println("\n注销服务请执行:\ncurl -X PUT -s http://127.0.0.1:8500/v1/agent/service/deregister/服务ID\n")

}

//编译: cd /home/icksys/tools/go/deleteService/consul && go build -o deleteService main.go
//调用示例: go run .\main.go -s auth3-server
