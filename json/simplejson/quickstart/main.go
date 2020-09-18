package main

import (
	"fmt"
	simplejson "github.com/bitly/go-simplejson"
)

func main() {
	//初始化*simpleJson.Json对象
	jsonStr := "{\"Name\":{\"en\":\"xiaobing\"},\"Age\":500,\"Birthday\":\"2011-11-11\",\"Sal\":8000,\"Skill\":\"牛魔拳\"}"
	sj, err := simplejson.NewJson([]byte(jsonStr))
	if err != nil {
		fmt.Printf("解析出错:%s\n", err.Error())
	}

	var v *simplejson.Json

	//获取字段，如果有多级，可以层层嵌套获取  	v = sj.Get("Name").Get(字段名2)
	v = sj.Get("Name").Get("en")

	//将v的值转化为具体类型的值，MustXXX方法一定可以转化成功
	//若转化不成功，则转化为该类型的零值
	result := v.MustString()
	fmt.Printf("解析结果:%s\n", result)

}

//func case2() {
//	//检查某个字段是否存在
//	_, ok := js.Get("字段名1").CheckGet("字段名2")
//	if ok {
//		fmt.Println("存在！")
//	} else {
//		fmt.Println("不存在")
//	}
//
//}
