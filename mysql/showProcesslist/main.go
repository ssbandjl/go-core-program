package main

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

	//二位数组
	cols, _ := rows.Columns() //列名
	var ret []map[string]string

	for rows.Next() {
		buff := make([]interface{}, len(cols))
		data := make([]string, len(cols))
		for i, _ := range buff {
			buff[i] = &data[i]
		}

		fmt.Printf("buff扫描前:\n%+v\n", buff)
		err := rows.Scan(buff...) //扫描到buff接口中，实际是字符串类型data中
		if err != nil {
			panic(err)
		}
		fmt.Printf("buff扫描后:\n%+v\n", buff)

		//将每一行数据存放到数组中
		dataKv := make(map[string]string, len(cols))
		fmt.Printf("打印列值:\n")
		for k, col := range data { //k是index，col是对应的值
			fmt.Printf("%30s:\t%s\n", cols[k], col)
			dataKv[cols[k]] = col
		}
		ret = append(ret, dataKv)
	}

	//fmt.Printf("执行结果:%+v\n", dataKv)
	//fmt.Printf("执行结果:%+v\n", ret[0])
	for k, v := range ret {
		fmt.Printf("执行结果:%+v %+v\n", k, v)

	}
	//fmt.Printf("Slave_IO_Running:%+v\n", dataKv["Slave_IO_Running"])
	//fmt.Printf("Slave_SQL_Running:%+v\n", dataKv["Slave_SQL_Running"])
	//fmt.Printf("Seconds_Behind_Master:%+v\n", dataKv["Seconds_Behind_Master"])

}
