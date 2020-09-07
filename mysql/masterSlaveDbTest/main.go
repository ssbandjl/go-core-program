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

type ShowMasterStatus struct {
	File            string
	Position        int
	BinlogDoDB      string
	BinlogIgnoreDB  string
	ExecutedGtidSet string
}

var (
	File            string
	Position        string
	BinlogDoDB      string
	BinlogIgnoreDB  string
	ExecutedGtidSet string
)

func GetDb(host string, port string, username string, password string) (*sql.DB, error) {

	//参数dataSourceName格式:用户名:密码@[tcp(localhost:3306)]/数据库名, 这里的Db只是句柄，使用的时候才会调用
	//mysqlConnStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/mysql?charset=utf8&parseTime=True&loc=Local", username, password, host, port)

	//mysqlConnStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/mysql?charset=utf8&parseTime=True&loc=Local", username, password, host, port)
	mysqlConnStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/mysql?charset=utf8&parseTime=True&loc=Local", username, password, host, port)
	//fmt.Printf("GetDb mysqlConnStr:%v\n", mysqlConnStr)
	Db, err = sql.Open("mysql", mysqlConnStr)
	//if err != nil {
	//	panic(err.Error())
	//}
	return Db, err
}

func main() {
	//获取数据库连接
	Db, err := GetDb("172.16.24.200", "31321", "root", "test123")
	if err != nil {
		panic(err)
	}

	//检查主库状态,记录File,Position信息
	rows, err2 := Db.Query("show master status")
	if err2 != nil {
		panic(err2)
	}

	defer rows.Close()

	//fmt.Printf("查询得到的rows:\n%+v",rows)
	//rows, _ := Db.Query("SHOW VARIABLES LIKE '%data%'")

	var showMasterStatus ShowMasterStatus
	var showMasterStatusList []ShowMasterStatus
	fmt.Println("开始扫描")
	//rows.Scan(&File, &Position)
	//fmt.Println(File, Position)
	//fmt.Printf("\n\n%v, %v, %v, %v, %v\n\n\n", File, Position, BinlogDoDB, BinlogIgnoreDB, ExecutedGtidSet)

	for rows.Next() {
		fmt.Printf("\n开始遍历\n")

		err := rows.Scan(&showMasterStatus.File, &showMasterStatus.Position, &showMasterStatus.BinlogDoDB, &showMasterStatus.BinlogIgnoreDB, &showMasterStatus.ExecutedGtidSet)
		//err:=rows.Scan(&File, &Position, &BinlogDoDB, &BinlogIgnoreDB)
		if err != nil {
			fmt.Printf("扫描数据字段出错:\n%s", err.Error())
		}
		//fmt.Printf("%v, %v, %v, %v, %v\n", File, Position, BinlogDoDB, BinlogIgnoreDB, ExecutedGtidSet)

		showMasterStatusList = append(showMasterStatusList, showMasterStatus)
		//fmt.Printf("主库状态信息:\n%+v", showMasterStatus)
	}
	fmt.Println(showMasterStatusList)
}

//
//
//	if rows.Next(){ //查询结果不为空
//		fmt.Printf("\n查询结果不为空\n")
//		//fmt.Printf("查询得到的rows:\n%+v",rows)
//		//rows.Scan(&File, &Position)
//		//fmt.Printf("%v, %v, %v, %v, %v\n", File, Position, BinlogDoDB, BinlogIgnoreDB, ExecutedGtidSet)
//
//
//
//		err=rows.Scan(&showMasterStatus.File, &showMasterStatus.Position, &showMasterStatus.BinlogDoDB, &showMasterStatus.BinlogIgnoreDB, &showMasterStatus.ExecutedGtidSet)
//		//err:=rows.Scan(&File, &Position, &BinlogDoDB, &BinlogIgnoreDB)
//		if err!=nil{
//			fmt.Printf("扫描数据字段出错:\n%s", err.Error())
//		}
//
//		fmt.Printf("主库状态信息:\n%+v", showMasterStatus)
//
//
//
//		for rows.Next(){
//			fmt.Printf("\n开始遍历\n")
//
//			err:=rows.Scan(&showMasterStatus.File, &showMasterStatus.Position, &showMasterStatus.BinlogDoDB, &showMasterStatus.BinlogIgnoreDB, &showMasterStatus.ExecutedGtidSet)
//			//err:=rows.Scan(&File, &Position, &BinlogDoDB, &BinlogIgnoreDB)
//			if err!=nil{
//				fmt.Printf("扫描数据字段出错:\n%s", err.Error())
//			}
//			//fmt.Printf("%v, %v, %v, %v, %v\n", File, Position, BinlogDoDB, BinlogIgnoreDB, ExecutedGtidSet)
//
//			showMasterStatusList=append(showMasterStatusList, showMasterStatus)
//			fmt.Printf("主库状态信息:\n%+v", showMasterStatus)
//			break
//		}
//	}else{
//		fmt.Printf("查询结果为空")
//	}
//	fmt.Sprintf("主库状态信息:\n%+v", showMasterStatusList)
//
//}
