package main

import (
	"context"
	"fmt"
)

//现在需要找到源文件中的这些地方：特征是调用了context.WithCancel函数，并且入参为nil。比如example/test2.go文件里面，有十多种可能
//从000到ccc，对应golang的AST的不同结构类型，现在需要把他们全部找出来。其中bbb这种情况代表了for语句，只不过在context.WithCancel函数不适用，所以注掉了。为了解决这个问题，首先需要仔细分析go/ast的Node接口
func test2(a string, b int) {
	context.WithCancel(nil) //000

	if _, err := context.WithCancel(nil); err != nil { //111
		context.WithCancel(nil) //222
	} else {
		context.WithCancel(nil) //333
	}

	_, _ = context.WithCancel(nil) //444

	go context.WithCancel(nil) //555

	go func() {
		context.WithCancel(nil) //666
	}()

	defer context.WithCancel(nil) //777

	defer func() {
		context.WithCancel(nil) //888
	}()

	data := map[string]interface{}{
		"x2": context.WithValue(nil, "k", "v"), //999
	}
	fmt.Println(data)

	/*
	   for i := context.WithCancel(nil); i; i = false {//aaa
	       context.WithCancel(nil)//bbb
	   }
	*/

	var keys []string = []string{"ccc"}
	for _, k := range keys {
		fmt.Println(k)
		context.WithCancel(nil)
	}
}
