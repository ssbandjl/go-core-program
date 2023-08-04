package main

import (
	"log"
	"os"
	"time"

	zk "github.com/samuel/go-zookeeper/zk"
)

var ZkConnected = true
var firstRegister bool = true

func dealErr(err error) {
	if err != nil {
		log.Printf("err:%s", err.Error())
		os.Exit(1)
	}
}

func main() {
	zkips := []string{"172.31.32.30:2181", "172.31.32.31:2181", "172.31.32.32:2181"}
	conn, session, err := zk.Connect(zkips, time.Duration(2)*time.Second)
	dealErr(err)
	chnQuit := make(chan bool)
	go func() {
		for {
			select {
			case <-chnQuit:
				log.Printf("Zookeeper event receiver quit now: session ID:%#x", conn.SessionID())
				return
			default:
				event, ok := <-session
				if !ok {
					log.Printf("Zookeeper error, fail to get session event session ID:%#x", conn.SessionID())
					time.Sleep(100 * time.Millisecond)
					continue
				}
				log.Printf("Zookeeper get a event, Type: %s, State:%s, path:%s, Server:%s, Error:%v, session ID:%#x", event.Type.String(),
					event.State.String(), event.Path, event.Server, event.Err, conn.SessionID())
				if event.State == zk.StateConnected {
					ZkConnected = true
					log.Printf("Zookeeper connected to:%s, session ID:%#x", conn.Server(), conn.SessionID())
				}
				if event.State == zk.StateDisconnected {
					ZkConnected = false
				}
				if event.State == zk.StateExpired {
					log.Printf("zk timeout exec sessionExpireHandler()")
				}
				if event.State == zk.StateHasSession {
					log.Printf("exec ZkInstance.sessionCreatedHandler()")

					{
						ParentPath := "/neonsan/neonsan0/centers"
						centerip := "172.31.32.30"
						children, _, err := conn.Children(ParentPath)
						c := conn
						dealErr(err)
						log.Printf("children:%s", children)
						for _, node := range children {
							childPath := ParentPath + "/" + node
							ip, stat, err := c.Get(childPath)
							dealErr(err)
							if string(ip) == centerip {
								log.Printf("stat.EphemeralOwner:%#x", stat.EphemeralOwner) // #x 十六进制表示，字母形式为小写 a-f, 如: 0x20422753814ae1c
								if firstRegister || stat.EphemeralOwner != c.SessionID() {
									log.Printf("zk path:%s already exist(owner session:%#x, current session:%#x), now delete it", childPath, stat.EphemeralOwner, c.SessionID())
									log.Printf("c.Delete(childPath, 0):%s", childPath)
									// if err = c.Delete(childPath, 0); err != nil {
									// 	log.Printf("Failed delete zk path:%s, error:%v", childPath, err)
									// }

								}
								log.Printf("center node: %s already created by zk session:%#x, current session:%#x", childPath, stat.EphemeralOwner, c.SessionID())
							}
						}
					}

				}
			}
		}
	}()

	time.Sleep(1000000000 * time.Second)
}
