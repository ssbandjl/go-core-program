package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go_code/util"
	"log"
	"time"
)

const (
	AutoDiscoverMongoPort = 27017
)

type MongoObj struct {
	Host   string
	Client *mongo.Client
}

func (this *MongoObj) GetClient() error {
	uri := fmt.Sprintf("mongodb://%s:%d", this.Host, AutoDiscoverMongoPort)
	clientOptions := options.Client().ApplyURI(uri)
	clientOptions.SetConnectTimeout(time.Second * 3)         //设置连接超时时间, 创建连接到服务器间的超时时间 ,参考:https://github.com/mongodb/mongo-go-driver/blob/master/mongo/options/clientoptions.go
	clientOptions.SetSocketTimeout(time.Second * 3)          //设置套接字超时时间, 即驱动等待套接字可以读取的时间
	clientOptions.SetServerSelectionTimeout(time.Second * 3) //设置服务器可用超时时间, 即驱动执行find命名的超时时间

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	//defer client.Disconnect(context.TODO())
	if err != nil {
		log.Printf("连接MongoDB失败:%s", err.Error())
		return err
	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Printf("连通性检查失败:%s", err.Error())
		return err
	}
	//util.Log4Zap(zap.InfoLevel).Info(fmt.Sprintf("连通性检查成功"))
	this.Client = client
	return nil
}

func (this *MongoObj) IsMaster() {
	cmd := bson.D{{"isMaster", "1"}}
	//cmd := bson.D{{"explain", kv}}
	opts := options.RunCmd().SetReadPreference(readpref.Primary())
	var result bson.M
	if err := this.Client.Database("admin").RunCommand(context.TODO(), cmd, opts).Decode(&result); err != nil {
		log.Printf("检查主节点失败:%s", err.Error())
	}
	log.Printf("检查主节点结果JSON:\n%s", util.Data2Json(result))
	log.Printf("检查主节点结果对象:\n%+v", result)
	//对象转map
	log.Printf("是否为主节点:%v", result["ismaster"])
}

func (this *MongoObj) RepSetStatus() {
	cmd := bson.D{{"serverStatus", "1"}}
	//cmd := bson.D{{"explain", kv}}
	opts := options.RunCmd().SetReadPreference(readpref.Primary())
	var result bson.M
	if err := this.Client.Database("admin").RunCommand(context.TODO(), cmd, opts).Decode(&result); err != nil {
		log.Printf("执行命令失败:%s", err.Error())
	}
	log.Printf("执行结果:\n%+v", result)
	//log.Printf("执行结果:\n%s", util.Data2Json(result))
}

func main() {
	mongoObj := MongoObj{Host: "data"}
	err := mongoObj.GetClient()
	if err != nil {
		log.Printf("获取Mongo客户端失败:%s", err.Error())
	}
	mongoObj.IsMaster()
	//mongoObj.RepSetStatus()
}
