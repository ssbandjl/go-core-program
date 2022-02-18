package main

import (
	"fmt"
	"log"
)

const (
	//MEM_SYNC_TYPE_VOLUME  event type of volume sync
	MEM_SYNC_TYPE_VOLUME uint16 = 1
	//MEM_SYNC_TYPE_STORE_STATISTICS event type of store statistics sync
	MEM_SYNC_TYPE_STORE_STATISTICS uint16 = 1 << 1
	//MEM_SYNC_TYPE_SSD_STATISTICS  event type of ssd statistics info sync
	MEM_SYNC_TYPE_SSD_STATISTICS uint16 = 1 << 2
)

var memEventCtrlMap map[uint16]string

func main() {
	memEventCtrlMap = make(map[uint16]string)
	memEventCtrlMap[MEM_SYNC_TYPE_VOLUME] = "MEM_SYNC_TYPE_VOLUME"
	memEventCtrlMap[MEM_SYNC_TYPE_STORE_STATISTICS] = "MEM_SYNC_TYPE_STORE_STATISTICS"
	memEventCtrlMap[MEM_SYNC_TYPE_SSD_STATISTICS] = "MEM_SYNC_TYPE_SSD_STATISTICS"
	fmt.Println(memEventCtrlMap)
	fmt.Println(MEM_SYNC_TYPE_VOLUME | MEM_SYNC_TYPE_STORE_STATISTICS)
	event_value := MEM_SYNC_TYPE_VOLUME | MEM_SYNC_TYPE_STORE_STATISTICS
	for event_type := range memEventCtrlMap {
		// log.Printf("event_value:%v, event_type:%v, event_value&event_type:%v", event_value, event_type, event_value&event_type)
		if (event_value & event_type) == event_type {
			log.Printf("event_value:%v, event_type:%v, event_value&event_type:%v,%s", event_value, event_type, event_value&event_type, memEventCtrlMap[event_type])
		}
	}
}
