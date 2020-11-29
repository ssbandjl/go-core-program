// +build go1.8

package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		c.String(http.StatusOK, "Welcome Gin Server")
	})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	// 用协程初始化一个服务, 它不会阻塞下面的优雅逻辑处理
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	//等待一个操作系统的中断信号, 来优雅的关闭服务
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM  //kill会发送终止信号
	// kill -2 is syscall.SIGINT  //发送强制进程结束信号
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it  //发送SIGKILL信号给进程
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit //阻塞在这里,直到获取到一个上面的信号
	log.Println("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	//这里使用context上下文包, 有5秒钟的处理超时时间
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil { //利用内置Shutdown方法优雅关闭服务
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
