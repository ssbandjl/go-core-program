package util

import (
	"encoding/json"
	"log"
)

func Data2Json(anyData interface{}) string {
	JsonByte, err := json.Marshal(anyData)
	if err != nil {
		log.Printf("数据序列化为json出错:\n%s\n", err.Error())
	}
	return string(JsonByte)
}
