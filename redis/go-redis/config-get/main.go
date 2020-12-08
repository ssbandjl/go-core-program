package main

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
)

var ctx = context.Background()

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "data:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	//检查连通性
	pingResult, err := client.Ping(ctx).Result()
	if err != nil {
		log.Printf("Redis连不通, 错误详情:\n%s", err)
	}
	log.Printf("Ping返回:%s", pingResult)

	//获取配置值
	val, err := client.ConfigGet(ctx, "appendfilename").Result()
	log.Printf("返回数组长度:%d", len(val))
	if err != nil {
		panic(err)
	}
	for _, valItem := range val {
		log.Printf("%s", valItem)
	}

}
