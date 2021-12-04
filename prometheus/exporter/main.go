package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Counter类型的Metric
var httpRequestCount = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_request_count",  // Metric的name
		Help: "http request count"}, // Metric的说明信息
	[]string{"endpoint"}) // Metric有一个Label，名称是endpoint，Metric形如 http_request_count(endpoint="")

// Gauge类型的Metric
var orderNum = prometheus.NewGauge(
	prometheus.GaugeOpts{
		Name: "order_num",
		Help: "order num"})

// Summary类型的Metric
var httpRequestDuration = prometheus.NewSummaryVec(
	prometheus.SummaryOpts{
		Name: "http_request_duration",
		Help: "http request duration",
	},
	[]string{"endpoint"},
)

// 将Metric注册到本地的Prometheus
func init() {
	prometheus.MustRegister(httpRequestCount)
	prometheus.MustRegister(orderNum)
	prometheus.MustRegister(httpRequestDuration)
}

func main() {
	// Exporter
	http.Handle("/metrics", promhttp.Handler()) // 对外暴露metrics接口，等待Prometheus来拉取
	http.HandleFunc("/hello/", hello)           // 处理业务请求，并变更Metric信息
	ipport := "127.0.0.1:8888"
	fmt.Println("服务器启动%s", ipport)
	err := http.ListenAndServe(ipport, nil)
	if err != nil {
		fmt.Println(err)
	}
}

// hello用于改变指标
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("process one request = %s\n", r.URL.Path)
	// Counter类型的Metric只能增
	// @Param lvs 表示label values
	httpRequestCount.WithLabelValues(r.URL.Path).Inc()
	start := time.Now()
	n := rand.Intn(100)
	// Gauge类型的Metric可增可减
	if n >= 90 {
		orderNum.Dec()
		time.Sleep(100 * time.Millisecond)
	} else {
		orderNum.Inc()
		time.Sleep(50 * time.Millisecond)
	}
	// Summary类型Metric
	elapsed := (float64)(time.Since(start) / time.Millisecond)
	httpRequestDuration.WithLabelValues(r.URL.Path).Observe(elapsed)
	w.Write([]byte("ok"))
}
