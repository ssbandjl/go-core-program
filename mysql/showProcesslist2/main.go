package main

//自动匹配字段
//上面查询的例子中，我们都自己定义了变量，同时查询的时候也写明了字段，如果不指名字段，或者字段的顺序和查询的不一样，都有可能出错。因此如果能够自动匹配查询的字段值，将会十分节省代码，同时也易于维护。
//go提供了Columns方法用获取字段名，与大多数函数一样，读取失败将会返回一个err，因此需要检查错误。
//代码例子如下：

//因为查询的时候是语句是：
//SELECT * FROM user_info WHERE user_id>6
//这样就会获取每行数据的所有的字段
//使用rows.Columns()获取字段名，是一个string的数组
//然后创建一个切片vals，用来存放所取出来的数据结果，类似是byte的切片。接下来还需要定义一个切片，这个切片用来scan，将数据库的值复制到给它
//vals则得到了scan复制给他的值，因为是byte的切片，因此在循环一次，将其转换成string即可。
//转换后的row即我们取出的数据行值，最后组装到result切片中。

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var (
	Db  *sql.DB
	err error
)

func GetDb(host string, port string, username string, password string) (*sql.DB, error) {
	mysqlConnStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/mysql?charset=utf8&parseTime=True&loc=Local", username, password, host, port)
	Db, err = sql.Open("mysql", mysqlConnStr)
	return Db, err
}

func main() {
	//获取数据库连接
	//Db, err := GetDb("172.16.24.200", "31422", "root", "test123")
	Db, err := GetDb("172.16.24.200", "31321", "root", "test123")
	if err != nil {
		panic(err)
	}
	rows, _ := Db.Query("show processlist")
	defer rows.Close()

	if err != nil {
		fmt.Println("select fail,err:", err)
		return
	}
	cols, err := rows.Columns()
	if err != nil {
		fmt.Println("get columns fail,err:", err)
		return
	}
	fmt.Println(cols)
	vals := make([][]byte, len(cols))
	scans := make([]interface{}, len(cols))

	for i := range vals {
		scans[i] = &vals[i]
	}
	fmt.Println(scans)
	var results []map[string]string

	for rows.Next() {
		err = rows.Scan(scans...)
		if err != nil {
			fmt.Println("scan fail,err:", err)
			return
		}
		row := make(map[string]string)
		for k, v := range vals {
			key := cols[k]
			row[key] = string(v)
		}
		results = append(results, row)
	}

	for k, v := range results {
		fmt.Println(k, v)
	}

}
