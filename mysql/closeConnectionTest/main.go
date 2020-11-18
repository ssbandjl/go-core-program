package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql" //被database/sql包使用的MYSQL驱动, 官方文档和项目:https://github.com/go-sql-driver/mysql
	"log"
	"time"
)

var (
	Db  *sql.DB
	err error
)

//获取数据库控制器
func GetDb(host string, port int, username string, password string) (*sql.DB, error) {
	//DSN (Data Source Name)数据源连接格式:[username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
	//这里我们可以不选择数据库,或者增加可选参数,比如timeout(建立连接超时时间)
	//mysqlConnStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/mysql?&charset=utf8&parseTime=True&loc=Local&timeout=5s", username, password, host, port)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/?timeout=5s", username, password, host, port)
	Db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Printf("配置连接出错:%s\n", err.Error())
		return Db, err
	}
	// 设置连接池中空闲连接的最大数量。
	Db.SetMaxIdleConns(1)
	// 设置打开数据库连接的最大数量。
	Db.SetMaxOpenConns(1)
	// 设置连接可复用的最大时间。
	Db.SetConnMaxLifetime(time.Second * 30)
	//设置连接最大空闲时间
	Db.SetConnMaxIdleTime(time.Second * 30)

	//检查连通性
	err = Db.Ping()
	if err != nil {
		log.Printf("数据库连接出错:%s\n", err.Error())
		return Db, err
	}

	return Db, err
}

//单行数据解析 查询数据库，解析查询结果，支持动态行数解析
func QueryAndParse(Db *sql.DB, queryStr string) map[string]string {
	rows, err := Db.Query(queryStr)
	defer rows.Close()

	if err != nil {
		log.Printf("查询出错,SQL语句:%s\n错误详情:%s\n", queryStr, err.Error())
		return nil
	}

	//获取列名cols
	cols, _ := rows.Columns()
	if len(cols) > 0 {
		buff := make([]interface{}, len(cols))       // 创建临时切片buff
		data := make([][]byte, len(cols))            // 创建存储数据的字节切片2维数组data
		dataKv := make(map[string]string, len(cols)) //创建dataKv, 键值对的map对象
		for i, _ := range buff {
			buff[i] = &data[i] //将字节切片地址赋值给临时切片,这样data才是真正存放数据
		}

		for rows.Next() {
			rows.Scan(buff...) // ...是必须的,表示切片
		}

		for k, col := range data {
			dataKv[cols[k]] = string(col)
			//fmt.Printf("%30s:\t%s\n", cols[k], col)
		}
		return dataKv
	} else {
		return nil
	}
}

//多行数据解析
func QueryAndParseRows(Db *sql.DB, queryStr string) []map[string]string {
	rows, err := Db.Query(queryStr)
	defer rows.Close()
	if err != nil {
		fmt.Printf("查询出错:\nSQL:\n%s, 错误详情:%s\n", queryStr, err.Error())
		return nil
	}
	//获取列名cols
	cols, _ := rows.Columns()
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
		return ret
	} else {
		return nil
	}
}

//任意可序列化数据转为Json,便于查看
func Data2Json(anyData interface{}) string {
	JsonByte, err := json.Marshal(anyData)
	if err != nil {
		log.Printf("数据序列化为json出错:\n%s\n", err.Error())
	}
	return string(JsonByte)
}

func main() {
	//获取数据库控制器DB
	DB, err := GetDb("data", 3306, "root", "root")
	if err != nil {
		log.Printf("获取数据库控制器出错:%s\n", err.Error())
	}
	defer DB.Close() //延迟关闭数据库控制器,释放数据库连接

	//单行数据查询
	showMasterStatus := QueryAndParse(Db, "show master status")
	log.Printf("单行数据-数据库状态:%v\n", Data2Json(showMasterStatus))
	log.Printf("单行数据-数据库状态-File:%v\n", showMasterStatus["File"])

	//多行数据查询
	showProcessList := QueryAndParseRows(Db, "show processlist")
	log.Printf("多行数据-进程信息:%v\n", Data2Json(showProcessList))
	log.Printf("多行数据-进程信息-Host:%v\n", showProcessList[0]["Host"])
}
