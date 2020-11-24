package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// Booking contains binded and validated data.
// Booking结构中定义了包含绑定器和日期验证器标签
type Booking struct {
	CheckIn  time.Time `form:"check_in" binding:"required,bookabledate" time_format:"2006-01-02"`     //登记时间
	CheckOut time.Time `form:"check_out" binding:"required,gtfield=CheckIn" time_format:"2006-01-02"` //gtfield=CheckIn表示结账时间必须大于登记时间
}

// 定义日期验证器
var bookableDate validator.Func = func(fl validator.FieldLevel) bool {
	date, ok := fl.Field().Interface().(time.Time) //利用反射获取到字段值 -> 转为接口 -> 类型断言(时间类型)
	if ok {
		today := time.Now()
		if today.After(date) { //如果当前时间在checkIn字段时间之后,返回false,即登记时间不能早于当前的时间
			return false
		}
	}
	return true
}

func main() {
	route := gin.Default()
	//对binding.Validator.Engine()接口进行类型断言,断言类型为Validate结构,如果断言成功,就将自定义的验证器注册到Gin内部
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// - if the key already exists, the previous validation function will be replaced. 该注册方法会将已经存在的验证器替换
		// - this method is not thread-safe it is intended that these all be registered prior to any validation
		// 注册方法不是线程安全的, 在验证开始前,需要保证所有的验证器都注册成功
		v.RegisterValidation("bookabledate", bookableDate)
	}

	route.GET("/bookable", getBookable)
	route.Run(":8085")
}

func getBookable(c *gin.Context) {
	var b Booking
	if err := c.ShouldBindWith(&b, binding.Query); err == nil {
		c.JSON(http.StatusOK, gin.H{"message": "Booking dates are valid!"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

//模拟请求:
// 登记时间和结账时间符合条件
//$ curl "localhost:8085/bookable?check_in=2030-04-16&check_out=2030-04-17"
//{"message":"Booking dates are valid!"}
//
// 登记时间在结账时间之后, 不满足gtfield校验规则
//$ curl "localhost:8085/bookable?check_in=2030-03-10&check_out=2030-03-09"
//{"error":"Key: 'Booking.CheckOut' Error:Field validation for 'CheckOut' failed on the 'gtfield' tag"}
//
// 登记时间在当前时间之前,不满足自定义的验证器
//$ curl "localhost:8085/bookable?check_in=2000-03-09&check_out=2000-03-10"
//{"error":"Key: 'Booking.CheckIn' Error:Field validation for 'CheckIn' failed on the 'bookabledate' tag"}%
