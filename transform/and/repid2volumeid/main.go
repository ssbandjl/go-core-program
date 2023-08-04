package main

import "log"

func ReplicaToVolumeId(repId uint64) uint64 {

	return repId &^ 0x0ffffff
}

func main() {
	log.Printf("volume id:%d", ReplicaToVolumeId(50331664)) //50331648
}
