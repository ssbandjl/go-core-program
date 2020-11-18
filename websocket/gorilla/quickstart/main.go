// websockets.go
package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

//协议升级器,将HTTP协议升级为websocket协议,简称ws或wss(https)
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	//允许跨域
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	//用http包新建/ws接口, 作为websocket服务端
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		//调用协议升级器的Upgrade方法与前端建立ws连接
		conn, err := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity
		if err != nil {
			fmt.Printf("[ERROR]:\n%s\n", err.Error())
			return

		}

		defer conn.Close()
		//向客户端发送消息类型为1(标准输出)的欢迎消息
		if err = conn.WriteMessage(1, []byte("连接到Websocket服务端成功\r\n")); err != nil {
			return
		}

		receiveMsg := ""
		for {
			// Read message from browser 从浏览器读取消息
			msgType, msg, err := conn.ReadMessage()
			if err != nil {
				fmt.Printf("[ERROR]从连接中读取消息失败,%+v\n", err)
				return
			}
			fmt.Printf("megType:%d, receiveMsg:%s\n", msgType, receiveMsg)
			receiveMsg = receiveMsg + string(msg) //拼接消息,直到遇到回车换行符"\r"
			fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))
			switch string(msg) {
			case "\r": //匹配换行
				receiveMsg = "\r\n您的输入是:" + receiveMsg + "\r\n"
				if err = conn.WriteMessage(msgType, []byte(receiveMsg)); err != nil {
					fmt.Printf("ERROR:%v\n", err)
					return
				}
				receiveMsg = ""
				continue
			}

			// Print the message to the console
			fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

			// Write message back to browser
			if err = conn.WriteMessage(msgType, msg); err != nil {
				fmt.Printf("ERROR:%v\n", err)
				return
			}
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "websockets.html")
	})
	fmt.Printf("websockets服务运行中\n")
	http.ListenAndServe(":8081", nil)
}
