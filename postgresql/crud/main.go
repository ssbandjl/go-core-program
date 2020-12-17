package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

const (
	// 	connStr := "postgres://postgres:cloud1688@172.16.24.200:31976?sslmode=disable&connect_timeout=3"
	// Initialize connection constants.
	HOST     = "172.16.24.200"
	PORT     = 31976
	DATABASE = "postgres"
	USER     = "postgres"
	PASSWORD = "cloud1688"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

type Db struct {
	db *sql.DB
}

// 创建表
func (this *Db) CreateTable() {
	// 以水果库存清单表inventory为例
	// Drop previous table of same name if one exists.  如果之前存在清单表, 则删除该表
	_, err := this.db.Exec("DROP TABLE IF EXISTS inventory;")
	checkError(err)
	fmt.Println("Finished dropping table (if existed)")

	// Create table. 创建表, 指定id, name, quantity(数量)字段, 其中id为主键
	_, err = this.db.Exec("CREATE TABLE inventory (id serial PRIMARY KEY, name VARCHAR(50), quantity INTEGER);")
	checkError(err)
	fmt.Println("Finished creating table")
}

// 删除表
func (this *Db) DropTable() {
	// 以水果库存清单表inventory为例
	// Drop previous table of same name if one exists.  如果之前存在清单表, 则删除该表
	_, err := this.db.Exec("DROP TABLE IF EXISTS inventory;")
	checkError(err)
	fmt.Println("Finished dropping table (if existed)")
}

// 增加数据
func (this *Db) Insert() {
	// Insert some data into table. 插入3条水果数据
	sql_statement := "INSERT INTO inventory (name, quantity) VALUES ($1, $2);"
	_, err := this.db.Exec(sql_statement, "banana", 150)
	checkError(err)
	_, err = this.db.Exec(sql_statement, "orange", 154)
	checkError(err)
	_, err = this.db.Exec(sql_statement, "apple", 100)
	checkError(err)
	fmt.Println("Inserted 3 rows of data")
}

// 读数据/查数据
func (this *Db) Read() {
	//读取数据
	// Read rows from table.
	var id int
	var name string
	var quantity int

	sql_statement := "SELECT * from inventory;"
	rows, err := this.db.Query(sql_statement)
	checkError(err)
	defer rows.Close()

	for rows.Next() {
		switch err := rows.Scan(&id, &name, &quantity); err {
		case sql.ErrNoRows:
			fmt.Println("No rows were returned")
		case nil:
			fmt.Printf("Data row = (%d, %s, %d)\n", id, name, quantity)
		default:
			checkError(err)
		}
	}
}

// 更新数据
func (this *Db) Update() {
	// Modify some data in table.
	sql_statement := "UPDATE inventory SET quantity = $2 WHERE name = $1;"
	_, err := this.db.Exec(sql_statement, "banana", 200)
	checkError(err)
	fmt.Println("Updated 1 row of data")
}

// 删除数据
func (this *Db) Delete() {
	// Delete some data from table.
	sql_statement := "DELETE FROM inventory WHERE name = $1;"
	_, err := this.db.Exec(sql_statement, "orange")
	checkError(err)
	fmt.Println("Deleted 1 row of data")
}

// 数据序列化为Json字符串, 便于人工查看
func Data2Json(anyData interface{}) string {
	JsonByte, err := json.Marshal(anyData)
	if err != nil {
		log.Printf("数据序列化为json出错:\n%s\n", err.Error())
		return ""
	}
	return string(JsonByte)
}

//多行数据解析
func QueryAndParseRows(Db *sql.DB, queryStr string) []map[string]string {
	rows, err := Db.Query(queryStr)
	defer rows.Close()
	if err != nil {
		log.Printf("查询出错:\nSQL:\n%s, 错误详情\n", queryStr, err.Error())
		return nil
	}
	cols, _ := rows.Columns() //列名
	if len(cols) > 0 {
		var ret []map[string]string //定义返回的映射切片变量ret
		for rows.Next() {
			buff := make([]interface{}, len(cols))
			data := make([][]byte, len(cols)) //数据库中的NULL值可以扫描到字节中
			for i, _ := range buff {
				buff[i] = &data[i]
			}
			rows.Scan(buff...) //扫描到buff接口中，实际是字符串类型data中

			//将每一行数据存放到数组中
			dataKv := make(map[string]string, len(cols))
			for k, col := range data { //k是index，col是对应的值
				//fmt.Printf("%30s:\t%s\n", cols[k], col)
				dataKv[cols[k]] = string(col)
			}
			ret = append(ret, dataKv)
		}
		log.Printf("返回多元素数组:\n%s", Data2Json(ret))
		return ret
	} else {
		return nil
	}
}

//单行数据解析 查询数据库，解析查询结果，支持动态行数解析
func QueryAndParse(Db *sql.DB, queryStr string) map[string]string {
	rows, err := Db.Query(queryStr)
	defer rows.Close()

	if err != nil {
		log.Printf("查询出错:\nSQL:\n%s, 错误详情\n", queryStr, err.Error())
		return nil
	}
	//rows, _ := Db.Query("SHOW VARIABLES LIKE '%data%'")

	cols, _ := rows.Columns()
	if len(cols) > 0 {
		buff := make([]interface{}, len(cols)) // 临时slice
		data := make([][]byte, len(cols))      // 存数据slice
		dataKv := make(map[string]string, len(cols))
		for i, _ := range buff {
			buff[i] = &data[i]
		}

		for rows.Next() {
			rows.Scan(buff...) // ...是必须的
		}

		for k, col := range data {
			dataKv[cols[k]] = string(col)
			//fmt.Printf("%30s:\t%s\n", cols[k], col)
		}
		log.Printf("返回单行数据Map:\n%s", Data2Json(dataKv))
		return dataKv
	} else {
		return nil
	}
}

func main() {
	// Initialize connection string. 初始化连接字符串, 参数包含主机,端口,用户名,密码,数据库名,SSL模式(禁用),超时时间
	var connectionString string = fmt.Sprintf("host=%s  port=%d user=%s password=%s dbname=%s sslmode=disable connect_timeout=3", HOST, PORT, USER, PASSWORD, DATABASE)

	// Initialize connection object. 初始化连接对象, 驱动名为postgres
	db, err := sql.Open("postgres", connectionString)
	defer db.Close()
	checkError(err)
	postgresDb := Db{
		db: db,
	}
	err = postgresDb.db.Ping() //连通性检查
	checkError(err)
	fmt.Println("Successfully created connection to database")

	postgresDb.CreateTable()                                     //创建表
	postgresDb.Insert()                                          //插入数据
	postgresDb.Read()                                            //查询数据
	QueryAndParseRows(postgresDb.db, "SELECT * from inventory;") //直接查询和解析多行数据
	QueryAndParse(postgresDb.db, "SHOW DateStyle;")              //直接查询和解析单行数据
	postgresDb.Update()                                          //修改/更新数据
	postgresDb.Read()
	postgresDb.Delete() //删除数据
	postgresDb.Read()
	postgresDb.DropTable()
}
