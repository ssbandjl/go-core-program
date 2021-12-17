package main

import (
	"math/rand"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/log"
)

type ClusterManager struct {
	Zone         string
	OOMCountDesc *prometheus.Desc //OOM错误计数
	RAMUsageDesc *prometheus.Desc //RAM使用指标
}

// Describe simply sends the two Descs in the struct to the channel.
// 实现Describe接口，传递指标描述符到channel
func (c *ClusterManager) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.OOMCountDesc
	ch <- c.RAMUsageDesc
}

// Collect函数将执行抓取函数并返回数据，返回的数据传递到channel中，并且传递的同时绑定原先的指标描述符。以及指标的类型（一个Counter和一个Guage）
func (c *ClusterManager) Collect(ch chan<- prometheus.Metric) {
	oomCountByHost, ramUsageByHost := c.ReallyExpensiveAssessmentOfTheSystemState()
	for host, oomCount := range oomCountByHost {
		ch <- prometheus.MustNewConstMetric(
			c.OOMCountDesc,          //描述符
			prometheus.CounterValue, //type
			float64(oomCount),       // value
			host,                    //label
		)
	}
	for host, ramUsage := range ramUsageByHost {
		ch <- prometheus.MustNewConstMetric(
			c.RAMUsageDesc,
			prometheus.GaugeValue,
			ramUsage,
			host,
		)
	}
}

// 创建结构体及对应的指标信息,NewDesc参数第一个为指标的名称，第二个为帮助信息，显示在指标的上面作为注释，第三个是定义的label名称数组，第四个是定义的Labels
func NewClusterManager(zone string) *ClusterManager {
	return &ClusterManager{
		Zone: zone,
		OOMCountDesc: prometheus.NewDesc(
			"clustermanager_oom_crashes_total", //指标名
			"Number of OOM crashes.",           //帮助信息
			[]string{"host"},                   //label名称数组
			prometheus.Labels{"zone": zone},    //labels
		),
		RAMUsageDesc: prometheus.NewDesc(
			"clustermanager_ram_usage_bytes",
			"RAM usage as reported to the cluster manager.",
			[]string{"host"},
			prometheus.Labels{"zone": zone},
		),
	}
}

// 真正昂贵的系统状态评估
func (c *ClusterManager) ReallyExpensiveAssessmentOfTheSystemState() (
	oomCountByHost map[string]int, ramUsageByHost map[string]float64,
) {
	oomCountByHost = map[string]int{
		"foo.example.org": int(rand.Int31n(1000)),
		"bar.example.org": int(rand.Int31n(1000)),
	}
	ramUsageByHost = map[string]float64{
		"foo.example.org": rand.Float64() * 100,
		"bar.example.org": rand.Float64() * 100,
	}
	return
}

func main() {
	workerDB := NewClusterManager("db")
	workerCA := NewClusterManager("ca")

	// Since we are dealing with custom Collector implementations, it might
	// be a good idea to try it out with a pedantic registry.
	reg := prometheus.NewPedanticRegistry()
	reg.MustRegister(workerDB)
	reg.MustRegister(workerCA)

	/*如果直接执行上面的参数的话，不会获取任何的参数，因为程序将自动推出，我们并未定义http接口去暴露数据出来，因此数据在执行的时候还需要定义一个httphandler来处理http请求。
	添加下面的代码到main函数后面，即可实现数据传递到http接口上：*/

	// 定义一个采集数据的收集器集合, 可以merge多个不同的采集数据到一个结果集合，这里我们传递了缺省的DefaultGatherer，所以他在输出中也会包含go运行时指标信息。同时包含reg是我们之前生成的一个注册对象，用来自定义采集数据
	gatherers := prometheus.Gatherers{
		prometheus.DefaultGatherer,
		reg,
	}

	/*promhttp.HandlerFor()函数传递之前的Gatherers对象，并返回一个httpHandler对象，这个httpHandler对象可以调用其自身的ServHTTP函数来接手http请求，并返回响应。其中promhttp.HandlerOpts定义了采集过程中如果发生错误时，继续采集其他的数据*/
	h := promhttp.HandlerFor(gatherers,
		promhttp.HandlerOpts{
			ErrorLog:      log.NewErrorLogger(),
			ErrorHandling: promhttp.ContinueOnError,
		})
	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	})
	log.Infoln("Start server at http://localhost:8080/metrics")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Errorf("Error occur when start server %v", err)
		os.Exit(1)
	}
}
