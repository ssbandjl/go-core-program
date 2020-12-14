package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	//connStr := "user=pqgotest dbname=pqgotest sslmode=verify-full"
	connStr := "postgres://postgres:cloud1688@172.16.24.200:31976?sslmode=disable&connect_timeout=3"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Printf("连接失败, 错误信息:%s", err.Error())
	}
	//age := 21
	//rows, err := db.Query("SELECT name FROM users WHERE age = $1", age)
	//…

	sql := "SELECT datname FROM pg_database;"
	rows, err := db.Query(sql)
	checkError(err)
	defer rows.Close()
	log.Printf("rows:%+v", rows)
	cols, _ := rows.Columns() //列名
	if len(cols) > 0 {
		var ret []map[string]string
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
		log.Printf("ret:%s", ret)
	}

	//for rows.Next() {
	//	switch err := rows.Scan(&id, &name, &quantity); err {
	//	case sql.ErrNoRows:
	//		fmt.Println("No rows were returned")
	//	case nil:
	//		fmt.Printf("Data row = (%d, %s, %d)\n", id, name, quantity)
	//	default:
	//		checkError(err)
	//	}
	//}
}
