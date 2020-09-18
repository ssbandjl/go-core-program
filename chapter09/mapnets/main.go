package main

import (
	"encoding/json"
	"fmt"
)

func Data2Json(anyData interface{}) string {
	JsonByte, err := json.Marshal(anyData)
	if err != nil {
		panic(err)
	}
	return string(JsonByte)
}

//存储主库标签信息
type MysqlInfo struct {
	Svc    string            `json:"svc"`
	Labels map[string]string `json:"labels"`
}

//参考数据: {"k8s集群IP":{"命名空间":{"master":{"svc":"master","labels":{"app":"master"}}, "slave":{"svc":"slave", "labels":{"app":"slave"}}}}}
type Store map[string]map[string]map[string]MysqlInfo

func main() {
	//master:=MysqlInfo{
	//	Svc: "master",
	//	Labels: map[string]string{"app":"master"},
	//}
	//
	//ns:=map[string]map[string]MysqlInfo{
	//	"命名空间": {
	//		"master": master,
	//	},
	//}
	//
	//store := Store{
	//	"k8s集群": ns,
	//}

	store := Store{
		"K8S集群": {
			"命名空间": {
				"master": MysqlInfo{
					Svc:    "master",
					Labels: map[string]string{"app": "test"}},
			},
		},
	}

	fmt.Printf("数据:%s", string(Data2Json(store)))

}
