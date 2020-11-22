package main

//Logger Middleware可以用于ELK收集metrics，例如，如果有ELK系统处理日志，则可以将Logger中间件中的Metrics以特定格式存放至特定Metrics日志文件中，后期给ELK分析
import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"time"
)

// 初始化Metrics日志
func InitMetrics(m Metric) error {
	if mlog == nil {
		mlog = logrus.New()
	}

	// 设置Rotate
	writer, err := rotatelogs.New(
		m.GetMetricPath()+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(m.GetMetricPath()),
		rotatelogs.WithRotationTime(METRICSROTATETIME),
		rotatelogs.WithRotationCount(METRICSROTATECOUNT),
	)
	if err != nil {
		return err
	}

	// 设置logger的Writer
	mlog.SetOutput(writer)
	// 设置日志格式为JSON，方便ELK分析
	mlog.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat:  "",
		DisableTimestamp: false,
		DataKey:          "",
		FieldMap:         nil,
		CallerPrettyfier: nil,
		PrettyPrint:      false,
	})

	return nil
}

// 自定义Metrics格式
type GinLogger struct {
	// StatusCode is HTTP response code.
	StatusCode int `json:"statuscode"`
	// Latency is how much time the server cost to process a certain request. (ms)
	Latency float64 `json:"latencyms"`
	// ClientIP equals Context's ClientIP method.
	ClientIP string `json:"clientip"`
	// Method is the HTTP method given to the request.
	Method string `json:"method"`
	// Path is a path the client requests.
	Path string `json:"path"`
	// BodySize is the size of the Response Body
	BodySize int `json:"bodysize"`

	// Common Fields
	// level
	Level string `json:"level"`
	// msg
	Msg string `json:"msg"`
	// time
	Time string `json:"time"`
	// Metrics Type
	MetricsType string `json:"metricstype"`
}

// 自定义给Gin的Formatter
var GinLoggerFormatter = func(param gin.LogFormatterParams) string {
	// 按需求定义哪些Metrics需要收集
	g := GinLogger{
		StatusCode:  param.StatusCode,
		Latency:     param.Latency.Seconds() * 1e3,
		ClientIP:    param.ClientIP,
		Method:      param.Method,
		Path:        param.Path,
		BodySize:    param.BodySize,
		Level:       "Info",
		Msg:         "RestRequest",
		Time:        time.Now().Format(time.RFC3339),
		MetricsType: TYPE_RESTHTTP_PERFORMANCE,
	}
	gjson, _ := json.Marshal(g)
	return string(gjson) + "\n"
}

// 设置给Gin Middleware的Output
func GetMetricsOutput() io.Writer {
	return mlog.Out
}

func main() {
	//r := gin.New()

	//默认Gin引擎Engine使用了Logger(记录请求信息)Recovery(从Panic中恢复)中间件: engine.Use(Logger(), Recovery())
	//Logger 支持记录一个API请求发生时间，返回Status Code，Latency(耗时)，远端IP，请求方法, 请求路径(URL Path)
	r := gin.Default()

	r.Use(Logger())

	r.GET("/ping", func(c *gin.Context) {
		example := c.MustGet("example").(string) //使用MustGet方法获取键值
		// 读取中间件在上下文中存储的内容
		log.Printf("中间件在上线文中存储的键值, key:example, value:%s", example)
		// 返回
		c.JSON(http.StatusOK, gin.H{"Response": "OK"})
	})

	r.Run(":8082")
}
