// advanced-middleware.go
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

//中间件Middleware是一个函数, 该函数签名表示中间件,传入一个http控制器方法(http.HandlerFunc), 返回的也是一个http控制器方法, 内部串联一系列操作(如增加日志打印), 达到包裹http控制器方法的作用
type Middleware func(http.HandlerFunc) http.HandlerFunc

// Logging logs all requests with its path and the time it took to process
// 打印所有请求requests的请求路径和请求时间
func Logging() Middleware {

	// Create a new Middleware
	// 创建并返回一个中间件, 这里使用的是匿名函数
	return func(f http.HandlerFunc) http.HandlerFunc {

		// Define the http.HandlerFunc
		// 定义中间件签名中要求返回的http控制器方法http.HandlerFunc
		return func(w http.ResponseWriter, r *http.Request) {

			// Do middleware things
			//该中间件具体要做的一系列操作
			start := time.Now()
			defer func() { log.Println(r.URL.Path, time.Since(start)) }()

			// Call the next middleware/handler in chain
			// 链式调用下一个中间件
			f(w, r)
		}
	}
}

// Method ensures that url can only be requested with a specific method, else returns a 400 Bad Request
func Method(m string) Middleware {

	// Create a new Middleware
	return func(f http.HandlerFunc) http.HandlerFunc {

		// Define the http.HandlerFunc
		return func(w http.ResponseWriter, r *http.Request) {

			// Do middleware things
			if r.Method != m {
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}

			// Call the next middleware/handler in chain
			f(w, r)
		}
	}
}

// Chain applies middlewares to a http.HandlerFunc
// 将中间件串联在http控制器方法后
func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}

func main() {
	http.HandleFunc("/", Chain(Hello, Method("GET"), Logging()))
	http.ListenAndServe(":8080", nil)
}

/*
$ go run advanced-middleware.go
2017/02/11 00:34:53 / 0s

$ curl -s http://localhost:8080/
hello world

$ curl -s -XPOST http://localhost:8080/
Bad Request
*/

func createNewMiddleware() Middleware {

	// Create a new Middleware
	middleware := func(next http.HandlerFunc) http.HandlerFunc {

		// Define the http.HandlerFunc which is called by the server eventually
		handler := func(w http.ResponseWriter, r *http.Request) {

			// ... do middleware things

			// Call the next middleware/handler in chain
			next(w, r)
		}

		// Return newly created handler
		return handler
	}

	// Return newly created middleware
	return middleware
}
