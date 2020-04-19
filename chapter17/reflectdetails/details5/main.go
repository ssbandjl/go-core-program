
package main
import (
	"reflect"
	"fmt"
)

//通过反射，修改,
// num int 的值
// 修改 student的值

func reflect01(b interface{}) {
	//2. 获取到 reflect.Value
	rVal := reflect.ValueOf(b)
	// 看看 rVal的Kind是 
	fmt.Printf("rVal kind=%v\n", rVal.Kind())
	//rVal kind=ptr

	//3. rVal
	//Elem返回v持有的接口保管的值的Value封装，或者v持有的指针指向的值的Value封装, 类似返回rVal指针对应的值*rVal
	rVal.Elem().SetInt(20)
}

func main() {

	var num int = 10
	reflect01(&num)
	fmt.Println("num=", num) // 20


	//你可以这样理解rVal.Elem()
	// num := 9
	// ptr *int = &num
	// num2 := *ptr  //=== 类似 rVal.Elem()
}