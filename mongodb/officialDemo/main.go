package main

import (
	"context"
	"fmt"
	"github.com/globalsign/mgo/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var client *mongo.Client

const (
	MongoDBUri   = "mongodb://data:27017"
	MyDataBase   = "officialDemo"
	MyCollection = "post"
)

//插入post
//The Go driver supports all of the newest features of MongoDB, including multi-document transactions, client-side encryption, bulk operations, and aggregation for advanced analytics cases. Working with MongoDB document data in Go is similar to working with JSON or XML
type Post struct {
	Title string `json:”title,omitempty”`
	Body  string `json:”body,omitempty”`
}

func InsertPost(title string, body string) {
	post := Post{title, body}
	//collection := client.Database(“my_database”).Collection(“posts”)
	collection := client.Database(MyDataBase).Collection(MyCollection)
	insertResult, err := collection.InsertOne(context.TODO(), post)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted post with ID:", insertResult.InsertedID)
}

//查询
func GetPost(id bson.ObjectId) {
	collection := client.Database(MyDataBase).Collection(MyCollection)
	filter := bson.D{{"id", id}}
	var post Post
	err := collection.FindOne(context.TODO(), filter).Decode(&post)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Found post with title", post.Title)

}

//官方最佳实践,参考链接:https://www.mongodb.com/golang
func main() {
	//连接MongoDB
	var err error
	client, err = mongo.NewClient(options.Client().ApplyURI(MongoDBUri))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("连接成功")
	defer client.Disconnect(ctx)

	//插入
	//InsertPost("测试标题", "测试主体")

	//查询

	GetPost(bson.ObjectId("5f8ffa0f704a06eaf65beba0"))

}
