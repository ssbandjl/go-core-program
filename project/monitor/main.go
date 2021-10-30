package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"math"
	"math/rand"
	"os"
	"path/filepath"
	"regexp"
	"time"
)

// 时区

var start, end, method string

type Metric struct {
	Ts    int64 `json:ts`
	Value int   `json:value`
}

func DealError(err error) {
	if err != nil {
		log.Printf(err.Error())
		panic(err)
	}
}

func RecordMetric() {
	rand.Seed(time.Now().UnixNano())
	for {
		todday := time.Now().Format("2006-01-02")
		metric := Metric{Ts: time.Now().UTC().Unix(), Value: rand.Intn(100)}
		// log.Printf("metric:%+v\n", metric)
		jsonData, err := json.Marshal(metric)
		DealError(err)
		log.Printf("%+s", jsonData)
		filePath := todday
		file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
		DealError(err)
		defer file.Close()
		writer := bufio.NewWriter(file)
		writer.WriteString(string(jsonData) + "\n")
		writer.Flush()
		time.Sleep(10 * time.Second)
	}
}

func GetMetric() {
	var Values []int
	// log.Printf("聚合方法:%s", method)
	startTimeObj, err := time.ParseInLocation("2006-01-02 15:04:05", start, time.Local) // https://www.jianshu.com/p/92d9344425a7
	DealError(err)
	endTimeObj, err := time.ParseInLocation("2006-01-02 15:04:05", end, time.Local)
	DealError(err)
	startTs := startTimeObj.Unix()
	endTs := endTimeObj.Unix()
	// fmt.Println(startTimeObj.Unix())
	// 遍历文件
	var files []string
	root := "./"
	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	DealError(err)
	for _, file := range files {
		_, err := time.Parse("2006-01-02", file)
		if err != nil {
			continue
		}
		fileNameTimeObj, _ := time.Parse("2006-01-02 15:04:05", fmt.Sprintf("%s 00:00:00", file))
		fileNameTimeTs := fileNameTimeObj.Unix()
		log.Printf("时间戳, start:%d, end:%d, file:%d", startTs, endTs, fileNameTimeTs)
		if (math.Abs(float64(startTs-fileNameTimeTs)) <= 3600*8) || startTs < fileNameTimeTs && fileNameTimeTs <= endTs {
			log.Printf("需要读取文件:%s", file)
			fileHandler, err := os.Open(file)
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
				var metric Metric
				err = json.Unmarshal([]byte(str), &metric)
				DealError(err)
				log.Printf("时间戳, start:%d,end:%d,metric:%d", startTs, endTs, metric.Ts)
				if startTs <= metric.Ts && metric.Ts <= endTs {
					Values = append(Values, metric.Value)
				}
			}
		}
	}
	// log.Printf("排序前values:%+v", Values)
	//排序
	QuickSort(0, len(Values)-1, Values)
	// log.Printf("排序后values:%+v", Values)

	switch method {
	case "min":
		log.Printf("最小值:%d", Values[0])
	case "max":
		log.Printf("最大值:%d", Values[len(Values)-1])
	case "avg":
		log.Printf("平均值:%d", SumArray(Values)/len(Values))
	}
}

func SumArray(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func QuickSort(left int, right int, array []int) {
	l := left
	r := right
	// pivot 是中轴， 支点,中间那个数
	pivot := array[(left+right)/2]
	// fmt.Printf("中轴pivot:%v\n", pivot)
	temp := 0

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

// go run main.go -start "2021-10-30 07:00:00" -end "2021-10-30 07:30:00" -method max
func main() {
	flag.StringVar(&start, "start", "", "开始时间2021-10-30 11:11:11")
	flag.StringVar(&end, "end", "", "结束时间2021-10-31 11:11:11")
	flag.StringVar(&method, "method", "", "聚合方法[min|max|avg]")
	flag.Parse()
	// log.Printf("start:%s, end:%s, method:%s", start, end, method)

	matchMethod, err := regexp.MatchString(`^min$|^max$|^avg$`, method)
	DealError(err)
	if matchMethod {
		GetMetric()
	} else {
		RecordMetric()

	}
}
