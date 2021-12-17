package main

import (
	"fmt"
	"net"
	"os"
	"time"

	"github.com/samuel/go-zookeeper/zk"
)

const (
	zkAddr = "192.168.7.240:2182"
)

func main() {
	go starServer("127.0.0.1:8897")
	go starServer("127.0.0.1:8898")
	go starServer("127.0.0.1:8899")

	a := make(chan bool, 1)
	<-a
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}

func starServer(port string) {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", port)
	fmt.Println("tcpAddr:", tcpAddr)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	//注册zk节点q
	// 链接zk
	conn, err := GetConnect()
	if err != nil {
		fmt.Printf(" connect zk error: %s ", err)
	}
	defer conn.Close()
	// zk节点注册
	err = RegistServer(conn, port)
	if err != nil {
		fmt.Printf(" regist node error: %s ", err)
	}

	for {
		time.Sleep(time.Second * 5)
		conn, err := listener.Accept()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s", err)
			continue
		}
		go handleCient(conn, port)
	}

	fmt.Println("aaaaaa")
}

func handleCient(conn net.Conn, port string) {
	defer conn.Close()

	daytime := time.Now().String()
	conn.Write([]byte(port + ": " + daytime))
}
func GetConnect() (conn *zk.Conn, err error) {
	zkList := []string{zkAddr}
	conn, _, err = zk.Connect(zkList, 10*time.Second)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func RegistServer(conn *zk.Conn, host string) (err error) {
	_, err = conn.Create("/go_servers/"+host, nil, zk.FlagEphemeral, zk.WorldACL(zk.PermAll))
	return
}

func GetServerList(conn *zk.Conn) (list []string, err error) {
	list, _, err = conn.Children("/go_servers")
	return
}
