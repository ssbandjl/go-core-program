package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(w, "HTTP, Hello")
	})

	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	http.HandleFunc("/ws", func(w http.ResponseWriter, req *http.Request) {
		conn, err := upgrader.Upgrade(w, req, nil)
		if err != nil {
			log.Println(err)
			return
		}
		defer conn.Close()

		//做一次读写
		var v interface{}
		err = conn.ReadJSON(&v)
		if err != nil {
			log.Println(err)
			return
		}
		log.Printf("接收到客户端:%s\n", v)
		if err := conn.WriteJSON("Hello WebSocket Lient"); err != nil {
			log.Println(err)
			return
		}
	})
	log.Fatal(http.ListenAndServe(":2021", nil))
}
