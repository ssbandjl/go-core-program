package main

import (
	"encoding/json"
	"fmt"
)

type Basic struct {
	Id       int     `json:"id"`
	Name     string  `json:"name"`
	Children []Basic `json:"children,omitempty"` //值为空时直接忽略,类似递归
}

func main() {
	//jsonData:=Instance{
	//	Basic: []Basic{
	//		{Id: 1, Name: "root", Children: []Basic{{Id: 2, Name: "master",}},}},
	//}
	//
	//str, err:=json.Marshal(jsonData)
	//if err !=nil{
	//	panic(err)
	//}
	//fmt.Println(string(str))

	jsonData := Basic{
		Id:   1,
		Name: "root",
		Children: []Basic{
			{Id: 2, Name: "master1", Children: []Basic{{Id: 3, Name: "slave1"}, {Id: 4, Name: "slave2"}}},
			{Id: 5, Name: "master2", Children: []Basic{{Id: 6, Name: "slave1"}, {Id: 7, Name: "slave2"}}},
		},
	}

	str, err := json.Marshal(jsonData)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(str))

}
