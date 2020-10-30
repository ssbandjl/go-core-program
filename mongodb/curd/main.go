package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

const (
	MongoDBUri = "mongodb://data:27017"
)

type Student struct {
	Name string
	Age  int
}

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
	collection := client.Database("demo").Collection("student")
	log.Printf("操作结果:%v", collection)

	////插入文档记录
	//s1 := Student{"小红", 12}
	//s2 := Student{"小兰", 10}
	//s3 := Student{"小黄", 11}
	//insertResult, err := collection.InsertOne(context.TODO(), s1)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	//
	//
	////插入多条记录
	//students := []interface{}{s2, s3}
	//insertManyResult, err := collection.InsertMany(context.TODO(), students)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)
	//
	//

	//更新文档, updateone()方法允许你更新单个文档。它需要一个筛选器文档来匹配数据库中的文档，并需要一个更新文档来描述更新操作。你可以使用bson.D类型来构建筛选文档和更新文档
	filter := bson.D{{"name", "小兰"}}
	//
	//update := bson.D{
	//	{"$inc", bson.D{
	//		{"age", 1},
	//	}},
	//}
	////增加1岁
	//updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)

	//查找文档, 要找到一个文档，你需要一个filter文档，以及一个指向可以将结果解码为其值的指针。要查找单个文档，使用collection.FindOne()。这个方法返回一个可以解码为值的结果。
	// 创建一个Student变量用来接收查询的结果
	var result Student
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("查找一个文档, Found a single document: %+v\n\n", result)

	// 查询多个
	// 将选项传递给Find(),只返回两个文档
	findOptions := options.Find()
	findOptions.SetLimit(2)

	// 定义一个切片用来存储查询结果
	var results []*Student

	// 把bson.D{{}}作为一个filter来匹配所有文档
	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	// 查找多个文档返回一个光标
	// 遍历游标允许我们一次解码一个文档
	for cur.Next(context.TODO()) {
		// 创建一个值，将单个文档解码为该值
		var elem Student
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// 完成后关闭游标
	cur.Close(context.TODO())
	fmt.Printf("查找多个文档, Found multiple documents (array of pointers): %#v\n\n", results)
	for _, v := range results {
		log.Printf("姓名:%s, 年龄:%d", v.Name, v.Age)
	}

	//删除文档, 可以使用collection.DeleteOne()或collection.DeleteMany()删除文档。如果你传递bson.D{{}}作为过滤器参数，它将匹配数据集中的所有文档。还可以使用collection. drop()删除整个数据集。
	// 删除名字是小黄的那个
	//deleteResult1, err := collection.DeleteOne(context.TODO(), bson.D{{"name","小黄"}})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("删除单条, Deleted %v documents in the trainers collection\n", deleteResult1.DeletedCount)
	//// 删除所有
	//deleteResult2, err := collection.DeleteMany(context.TODO(), bson.D{{}})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("删除全部文档, Deleted %v documents in the trainers collection\n", deleteResult2.DeletedCount)
	//

	// 断开连接
	err = client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("关闭连接, Connection to MongoDB closed.")

}
