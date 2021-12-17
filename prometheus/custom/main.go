package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Azure/go-amqp"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type ClusterManager struct {
	Zone string
	// Contains many more fields not listed in this example.
}

func (c *ClusterManager) ReallyExpensiveAssessmentOfTheSystemState() (
	oomCountByHost map[string]int, ramUsageByHost map[string]float64,
) {
	// Just example fake data.
	oomCountByHost = map[string]int{
		"foo.example.org": 42,
		"bar.example.org": 2001,
	}
	ramUsageByHost = map[string]float64{
		"foo.example.org": 6.023e23,
		"bar.example.org": 3.14,
	}
	return
}

type ClusterManagerCollector struct {
	ClusterManager *ClusterManager
}

var (
	oomCountDesc = prometheus.NewDesc(
		"clustermanager_oom_crashes_total",
		"Number of OOM crashes.",
		[]string{"host"}, nil,
	)
	ramUsageDesc = prometheus.NewDesc(
		"clustermanager_ram_usage_bytes",
		"RAM usage as reported to the cluster manager.",
		[]string{"host"}, nil,
	)
)

func (cc ClusterManagerCollector) Describe(ch chan<- *prometheus.Desc) {
	prometheus.DescribeByCollect(cc, ch)
}

func (cc ClusterManagerCollector) Collect(ch chan<- prometheus.Metric) {
	oomCountByHost, ramUsageByHost := cc.ClusterManager.ReallyExpensiveAssessmentOfTheSystemState()
	for host, oomCount := range oomCountByHost {
		ch <- prometheus.MustNewConstMetric(
			oomCountDesc,
			prometheus.CounterValue,
			float64(oomCount),
			host,
		)
	}
	for host, ramUsage := range ramUsageByHost {
		ch <- prometheus.MustNewConstMetric(
			ramUsageDesc,
			prometheus.GaugeValue,
			ramUsage,
			host,
		)
	}
}

func NewClusterManager(zone string, reg prometheus.Registerer) *ClusterManager {
	c := &ClusterManager{
		Zone: zone,
	}
	cc := ClusterManagerCollector{ClusterManager: c}
	prometheus.WrapRegistererWith(prometheus.Labels{"zone": zone}, reg).MustRegister(cc)
	return c
}

func failOnError(err error, input string) {
	log.Printf(err.Error(), input)
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare("hello", false, false, false, false, nil)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			var streams []byte
			streams = d.Body

			var metrics myMap
			err := json.Unmarshal(streams, &metrics)
			if err != nil {
				fmt.Println(err)
			}

			var category string
			category = metrics.Resource.Category

			if category == "server" {
				myMap := make(map[string]float64)
				MyMap[metrics.Resource.ResourceDataList[0].ResourceId] = metrics.Resource.ResourceDataList[0].MetricSampleList[0].ValueArray[0]
			}
		}
	}()

	reg := prometheus.NewPedanticRegistry()

	NewClusterManager("zone", reg)

	reg.MustRegister(
		prometheus.NewProcessCollector(prometheus.ProcessCollectorOpts{}),
		prometheus.NewGoCollector(),
	)

	http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{}))
	log.Fatal(http.ListenAndServe(":8080", nil))
	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever
}

//relevant structs go here for parsing JSON

// func (c *ClusterManager) ReallyExpensiveAssessmentOfTheSystemState() map[string]float64 {
// 	myMap := make(map[string]float64)
// 	myMap["device1"] = 754
// 	myMap["device2"] = 765

// 	return myMap
// }

// func (cc ClusterManagerCollector) Collect(ch chan<- prometheus.Metric) {
// 	values := cc.ClusterManager.ReallyExpensiveAssessmentOfTheSystemState()

// 	for key, value := range values {
// 		ch <- prometheus.MustNewConstMetric(
// 			valueDesc,
// 			prometheus.CounterValue,
// 			value,
// 			key,
// 		)
// 	}
// }
