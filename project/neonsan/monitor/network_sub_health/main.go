package main

import (
	"log"
	"sync"
	"time"
)

var (
	GlobalDataFromCenter *DataFromCenter
)

type DataFromCenter struct {
	StoreIps []string
	sync.Mutex
}

func NewDataFromCenter() *DataFromCenter {
	return &DataFromCenter{}
}

func (dfc *DataFromCenter) Update() (err error) {
	dfc.Lock()
	defer dfc.Unlock()
	var Stores = []string{"192.168.1.1", "192.168.1.2"}
	dfc.StoreIps = []string{}
	for _, ip := range Stores {
		dfc.StoreIps = append(dfc.StoreIps, ip)
	}
	return nil
}

func UpdateDataFromCenterLoop() {
	for {
		err := GlobalDataFromCenter.Update()
		if err != nil {
			log.Printf("err:%s", err.Error())
		}
		// log.Printf("%v", GlobalDataFromCenter)
		// time.Sleep(5 * time.Second)
	}
}

func main() {
	GlobalDataFromCenter = NewDataFromCenter()
	go UpdateDataFromCenterLoop()
	for {
		GlobalDataFromCenter.Lock()
		log.Printf("ip:%v", GlobalDataFromCenter.StoreIps)
		GlobalDataFromCenter.Unlock()

		// for _, ip := range GlobalDataFromCenter.StoreIps {
		// 	// log.Printf("ip:%s", ip)
		// }
	}

	time.Sleep(100000 * time.Second)
}
