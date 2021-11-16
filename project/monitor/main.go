package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
)

const (
	ScrapeInterval = 30
	MetricKeepDays = 3 //90
	MetricDataDir  = "data"
	Separator      = os.PathSeparator
)

var start, end, metric, method, labels string

type Metric struct {
	Ts     int64             `json:"ts"`
	Name   string            `json:"name"`
	Value  float64           `json:"value"`
	Labels map[string]string `json:"labels"`
}

func DealError(err error) {
	if err != nil {
		log.Printf(err.Error())
		panic(err)
	}
}

func CreateDir(dir string) error {
	_, err := os.Stat(dir)
	if err != nil {
		return os.MkdirAll(dir, os.ModePerm)
	} else {
		return nil
	}
}

func CheckMetricFile() {
	for {
		files, err := GetDirFiles(MetricDataDir)
		// log.Printf("指标数据文件列表:%+v, 个数:%d", files, len(files))
		DealError(err)
		if len(files) > MetricKeepDays {
			sort.Strings(files)
			log.Printf("指标数据文件个数超过配置值%d,删除最早的历史文件:%s", MetricKeepDays, files[0])
			err = os.Remove(files[0])
			DealError(err)
		}
		time.Sleep(3600 * 24 * time.Second)
	}
}

func RecordMetric() {
	CreateDir(MetricDataDir)
	rand.Seed(time.Now().UnixNano())
	metricMap := make(map[string]float64)
	for {
		today := time.Now().Format("2006-01-02")
		load, err := load.Avg()
		DealError(err)
		metricMap["load1"] = load.Load1
		metricMap["load5"] = load.Load5
		metricMap["load15"] = load.Load15

		cpuPercent, err := cpu.Percent(time.Second, false)
		DealError(err)
		metricMap["cpuPercent"] = cpuPercent[0]

		memInfo, err := mem.VirtualMemory()
		DealError(err)
		// log.Printf("memory:%+v", memInfo)
		metricMap["memoryTotal"] = float64(memInfo.Total)
		metricMap["memoryAvailable"] = float64(memInfo.Available)
		metricMap["memoryFree"] = float64(memInfo.Free)
		metricMap["memoryUsed"] = float64(memInfo.Used)
		metricMap["memoryBuffers"] = float64(memInfo.Buffers)
		metricMap["memoryCached"] = float64(memInfo.Cached)

		// parts, err := disk.Partitions(true)
		// DealError(err)
		// for _, part := range parts {
		// 	fmt.Printf("part:%v\n", part.String())
		// 	diskInfo, _ := disk.Usage(part.Mountpoint)
		// 	fmt.Printf("disk info:used:%v free:%v\n", diskInfo.UsedPercent, diskInfo.Free)
		// }

		// ioStat, _ := disk.IOCounters()
		// for k, v := range ioStat {
		// 	fmt.Printf("%v:%v\n", k, v)
		// }

		// netInfo, _ := net.IOCounters(true)
		// for index, v := range netInfo {
		// 	fmt.Printf("%v:%v send:%v recv:%v\n", index, v, v.BytesSent, v.BytesRecv)
		// }

		for k, v := range metricMap {
			// metric := Metric{Ts: time.Now().UTC().Unix(), Name:k, Value: float64(rand.Intn(100))}
			labels := map[string]string{"host": "localhost", "system": "linux"}
			metricObj := Metric{Ts: time.Now().UTC().Unix(), Name: k, Value: v, Labels: labels}
			// log.Printf("metric:%+v\n", metric)
			jsonData, err := json.Marshal(metricObj)
			DealError(err)
			log.Printf("json指标:%+s", jsonData)
			filePath := filepath.Join(MetricDataDir, today)
			file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
			DealError(err)
			defer file.Close()
			writer := bufio.NewWriter(file)
			writer.WriteString(string(jsonData) + "\n")
			writer.Flush()
		}
		time.Sleep(ScrapeInterval * time.Second)
	}
}

func GetDirFiles(dir string) (files []string, err error) {
	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if path != dir {
			files = append(files, path)
		}
		return nil
	})
	return
}

func GetMetric() float64 {
	var ret float64
	var Values []float64
	// log.Printf("聚合方法:%s", method)
	startTimeObj, err := time.ParseInLocation("2006-01-02 15:04:05", start, time.Local) // https://www.jianshu.com/p/92d9344425a7
	DealError(err)
	endTimeObj, err := time.ParseInLocation("2006-01-02 15:04:05", end, time.Local)
	DealError(err)
	startTs := startTimeObj.Unix()
	endTs := endTimeObj.Unix()
	files, err := GetDirFiles(MetricDataDir)
	DealError(err)
	for _, file := range files {
		fileName := strings.ReplaceAll(file, MetricDataDir+string(Separator), "")
		_, err := time.Parse("2006-01-02", fileName)
		if err != nil {
			// log.Printf("文件名非日期格式,%s", err.Error())
			continue
		}
		// fileNameTimeObj, _ := time.Parse("2006-01-02 15:04:05", fmt.Sprintf("%s 00:00:00", file))
		fileNameTimeObj, _ := time.ParseInLocation("2006-01-02", fileName, time.Local)
		fileNameTimeTs := fileNameTimeObj.Unix()
		// log.Printf("时间戳, start:%d, end:%d, file:%d", startTs, endTs, fileNameTimeTs)
		// if (math.Abs(float64(startTs-fileNameTimeTs)) <= 3600*8) || startTs < fileNameTimeTs && fileNameTimeTs <= endTs {
		if startTs <= (fileNameTimeTs+3600*24) && fileNameTimeTs <= endTs {
			filePath := filepath.Join(MetricDataDir, fileName)
			log.Printf("读取指标文件:%s", filePath)
			fileHandler, err := os.Open(filePath)
			DealError(err)
			defer fileHandler.Close()
			reader := bufio.NewReader(fileHandler)
			//循环的读取文件的内容
			for {
				str, err := reader.ReadString('\n') // 读到一个换行就结束
				if err == io.EOF {                  // io.EOF表示文件的末尾
					break
				}
				//输出内容
				// fmt.Printf(str)
				var metricObj Metric
				err = json.Unmarshal([]byte(str), &metricObj)
				DealError(err)
				// log.Printf("时间戳, start:%d,end:%d,metric:%d", startTs, endTs, metric.Ts)
				// 过滤

				if startTs <= metricObj.Ts && metricObj.Ts <= endTs && metricObj.Name == metric {
					labelMatch := true
					if labels != "" {
						labelStrs := strings.Split(labels, ",")
						for _, labelStr := range labelStrs {
							k := strings.Split(labelStr, "=")[0]
							v := strings.Split(labelStr, "=")[1]
							// log.Printf("labelsStr:%+v:%+v", k, v)
							if metricObj.Labels[k] != v {
								labelMatch = false
							}
						}
						if labelMatch {
							Values = append(Values, metricObj.Value)
						}
					} else {
						Values = append(Values, metricObj.Value)

					}
				}
			}
		}
	}

	// log.Printf("排序前values:%+v", Values)
	// 快速排序
	if len(Values) == 0 {
		log.Printf("not find metric")
		return ret
	}
	// QuickSort(0, len(Values)-1, Values)
	sort.Float64s(Values)
	// log.Printf("排序后values:%+v", Values)

	switch method {
	case "min":
		log.Printf("最小值:%f", Values[0])
		ret = Values[0]
	case "max":
		log.Printf("最大值:%f", Values[len(Values)-1])
		ret = Values[len(Values)-1]
	case "avg":
		log.Printf("平均值:%f", SumArray(Values)/float64(len(Values)))
		ret = SumArray(Values) / float64(len(Values))
	}
	return ret
}

func SumArray(array []float64) float64 {
	result := 0.0
	for _, v := range array {
		result += v
	}
	return result
}

func QuickSort(left int, right int, array []float64) {
	l := left
	r := right
	// pivot 是中轴， 支点,中间那个数
	pivot := array[(left+right)/2]
	// fmt.Printf("中轴pivot:%v\n", pivot)
	var temp float64

	//for 循环的目标是将比 pivot 小的数放到 左边
	//  比 pivot 大的数放到 右边
	for l < r { //退出条件就是l >= r
		//从  pivot 的左边找到大于等于pivot的值, 因为要交换, 修改这里的两个比较符号可以修改排序方向
		for array[l] < pivot {
			l++
		}
		//从  pivot 的右边边找到小于等于pivot的值
		for array[r] > pivot {
			r--
		}
		// 1 >= r 表明本次分解任务完成, break,说明未找到
		if l >= r {
			break
		}
		//交换
		temp = array[l]
		array[l] = array[r]
		array[r] = temp
		// fmt.Printf("交换后:%v\n", array)

		//优化，如果左/右指针走到中间位置,需要移动一位，便于递归
		if array[l] == pivot {
			r--
		}
		if array[r] == pivot {
			l++
		}
		// for循环结束后,l可能比r小
	}
	// 如果  1== r, 再移动下, 不要在比较了, 防止死循环
	if l == r {
		l++
		r--
	}
	// 向左递归
	// fmt.Println(l, r)
	if left < r {
		QuickSort(left, r, array)
	}
	// 向右递归
	if right > l {
		QuickSort(l, right, array)
	}
}

// go run main.go -start "2021-10-30 07:00:00" -end "2021-10-30 07:30:00" -metric load1 -method max -labels "host=localhost,system=linux"
func main() {
	if len(os.Args) != 1 && len(os.Args) < 9 {
		fmt.Println(`参数有误: 示例:go run main.go -start "2021-10-30 07:00:00" -end "2021-10-30 07:30:00" -metric load1 -method max -labels "host=localhost,system=linux"`)
		return
	}
	flag.StringVar(&start, "start", "", "开始时间2021-10-30 11:11:11")
	flag.StringVar(&end, "end", "", "结束时间2021-10-31 11:11:11")
	flag.StringVar(&metric, "metric", "load1", "指标名[load1|load5|load15]")
	flag.StringVar(&method, "method", "", "聚合方法[min|max|avg]")
	flag.StringVar(&labels, "labels", "", "标签[可选], -labels host=localhost,system=linux")
	flag.Parse()
	log.Printf("参数,start:%s, end:%s, metric:%s, method:%s, labels:%s", start, end, metric, method, labels)
	matchMethod, err := regexp.MatchString(`^min$|^max$|^avg$`, method)
	DealError(err)
	if matchMethod {
		go CheckMetricFile()
		GetMetric()
	} else {
		RecordMetric()
	}
}
