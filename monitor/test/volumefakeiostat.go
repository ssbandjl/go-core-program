package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

var Monitor_ip = flag.String("ip", "127.0.0.1", "the ip of monitor")
var Thread_count = flag.Int("threadNum", 10, "how much thread send data to monitor")
var Volume_count = flag.Int("volumeNum", 1000, "how much volumes' io data send per thread ")

type VolumeIostat struct {
	Volume_id       int    `json:"volume_id"`
	Volume_name     string `json:"volume_name"`
	Start_time      int64  `json:"start_time"`
	End_time        int64  `json:"end_time"`
	Read_iops       int64  `json:"read_iops"`
	Write_iops      int64  `json:"write_iops"`
	Read_latency    int64  `json:"read_latency"`
	Write_latency   int64  `json:"write_latency"`
	Read_bandwidth  int64  `json:"read_bandwidth"`
	Write_bandwidth int64  `json:"write_bandwidth"`
	Read_size       int64  `json:"read_size"`
	Write_size      int64  `json:"write_size"`
}

func Post(client *http.Client, url string, data []*VolumeIostat, contentType string) string {
	jsonStr, _ := json.Marshal(data)
	resp, err := client.Post(url, contentType, bytes.NewReader(jsonStr))
	if err != nil {
		fmt.Println(err)
		return ""
		// panic(err)
	}
	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)
	return string(result)
}

func running() {
	vlist := make([]*VolumeIostat, 0)
	var neonTransport http.RoundTripper = &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          100,
		MaxIdleConnsPerHost:   20,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
	client := &http.Client{Timeout: 5 * time.Second, Transport: neonTransport}
	var Ticker = time.NewTicker(time.Millisecond * 1000)
	for {
		if _, ok := <-Ticker.C; ok {
			End_time := time.Now().Unix()
			Start_time := End_time - 1
			for i := 1; i <= *Volume_count; i++ {
				tempv := &VolumeIostat{
					Volume_id:       i,
					Volume_name:     strconv.Itoa(i),
					Start_time:      Start_time,
					End_time:        End_time,
					Read_iops:       4096,
					Write_iops:      4096,
					Read_latency:    1,
					Write_latency:   1,
					Read_bandwidth:  4096000,
					Write_bandwidth: 4096000,
					Read_size:       4096,
					Write_size:      4096,
				}
				vlist = append(vlist, tempv)
			}
			// jsonStr, _ := json.Marshal(vlist)
			_ = Post(client, "http://"+*Monitor_ip+":2610/stat?op=report_volume_io", vlist, "application/json")
			// fmt.Println(res)
			vlist = make([]*VolumeIostat, 0)
			// client.CloseIdleConnections()
		}
	}
}
func main() {
	flag.Parse()
	for i := 0; i < *Thread_count; i++ {
		go running()
	}
	exitChan := make(chan int)
	signalChan := make(chan os.Signal, 1)
	go func() {
		<-signalChan
		fmt.Println("signal received")
		exitChan <- 1
	}()
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-exitChan
}
