package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	MongoDBUri = "mongodb://data:27017"
)

//连接池模式
func ConnectToDB(uri, name string, timeout time.Duration, num uint64) (*mongo.Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	o := options.Client().ApplyURI(uri)
	o.SetMaxPoolSize(num)
	client, err := mongo.Connect(ctx, o)
	if err != nil {
		return nil, err
	}

	return client.Database(name), nil
}

func main() {
	// 设置客户端连接配置
	//clientOptions := options.Client().ApplyURI(MongoDBUri)

	// 连接到MongoDB
	//client, err := mongo.Connect(context.TODO(), clientOptions)
	//if err != nil {
	//	log.Fatal(err)
	//}
	database, err := ConnectToDB(MongoDBUri, "demo", 10000, 5)

	//// 检查连接
	//err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	//log.Printf("%s, 连接成功", MongoDBUri)
	//fmt.Println("连接成功: Connected to MongoDB!")

	// 指定获取要操作的数据集
	//collection := client.Database("q1mi").Collection("student")
	//log.Printf("操作结果:%v", collection)

	//// 断开连接
	//err = client.Disconnect(context.TODO())
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println("Connection to MongoDB closed.")

}
