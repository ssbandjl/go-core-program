package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

// You will be using this Trainer type later in the program, 下面申明Trainer"教练员"结构体
type Trainer struct {
	Name string
	Age  int
	City string
}

func main() {
	// Rest of the code will go here  这里先占位

	// Set client options  设置客户端选项,如这里的连接字符串
	clientOptions := options.Client().ApplyURI("mongodb://data:27017")

	// Connect to MongoDB 连接数据库,返回客户端对象client
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection  连通性检查
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	collection := client.Database("test").Collection("trainers")

	//bson.D{{
	//	"name",
	//	bson.D{{
	//		"$in",
	//		bson.A{"Alice", "Bob"}
	//	}}
	//}}

	//插入文档,首先创建多个教练员结构体
	ash := Trainer{"Ash", 10, "Pallet Town"}
	misty := Trainer{"Misty", 10, "Cerulean City"}
	brock := Trainer{"Brock", 15, "Pewter City"}

	//使用InsertOne方法插入单个文档
	insertResult, err := collection.InsertOne(context.TODO(), ash)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)

	//使用InsertMany插入多个文档,该方法接收教练员切片作为参数
	trainers := []interface{}{misty, brock}

	insertManyResult, err := collection.InsertMany(context.TODO(), trainers)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)

	//更新文档
	//需要先构造一个匹配文档的D对象
	filter := bson.D{{"name", "Ash"}}

	//inc为递增,这里将名为Ash的教练员年龄+1
	update := bson.D{
		{"$inc", bson.D{
			{"age", 1},
		}},
	}

	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)

	//查找文档
	// create a value into which the result can be decoded 需要先申明一个教练员结构
	var result Trainer

	//FindOne方法接收一个过滤器对象,依然使用上面的Ash教练员,然后将查询结果解码到一个可解码的对象地址(指针类型)
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found a single document: %+v\n", result)

	//查找多个文档
	// Pass these options to the Find method
	findOptions := options.Find()
	findOptions.SetLimit(2) //设置查询结果限制为2个

	// Here's an array in which you can store the decoded documents 申明结果为一个教练员切片指针
	var results []*Trainer

	// Passing bson.D{{}} as the filter matches all documents in the collection, 查询后得到一个游标对象cur
	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time 遍历游标对象
	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem Trainer
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, &elem) //将单个查询结果添加到教练员切片中
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// Close the cursor once finished 游标用完后需要关闭,防止内存泄漏
	cur.Close(context.TODO())

	fmt.Printf("Found multiple documents (array of pointers): %+v\n", results)

	//删除文档
	//删除单个对象使用collection.DeleteOne(),如删除Ash教练员
	deleteOneResult, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteOneResult.DeletedCount)

	//DeleteMany删除多个文档,空D对象bson.D{{}}作为过滤参数表示删除所有文档集, 也可以使用collection.Drop()删除整个文档集合
	deleteResult, err := collection.DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)

	//最后记得关闭连接,防止内存泄漏
	err = client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")

}
