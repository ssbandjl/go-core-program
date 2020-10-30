package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	//MongoDBUri = "mongodb://data:27017"
	MongoDBUri = "mongodb://172.16.13.117:30942"
)

func main() {
	// 设置客户端连接配置
	clientOptions := options.Client().ApplyURI(MongoDBUri)

	// 连接到MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// 检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s, 连接成功", MongoDBUri)
	//fmt.Println("连接成功: Connected to MongoDB!")

	// 指定获取要操作的数据集
	//collection := client.Database("q1mi").Collection("student")
	//log.Printf("操作结果:%v", collection)

	// 断开连接
	err = client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")

}
