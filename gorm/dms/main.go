package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

//定义结构
type User struct {
	ID          int       `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	IsSuperuser bool      `gorm:"type:tinyint;column:is_superuser;not null;default:0;" json:"is_superuser"`
	IsActive    bool      `gorm:"type:tinyint;column:is_active;not null;default:1;" json:"is_active"`
	Username    string    `gorm:"type:varchar(128);column:username;unique_index;not null;" json:"username"`
	Password    string    `gorm:"type:varchar(256);column:password;" json:"password"`
	Mobile      string    `gorm:"type:varchar(32);column:mobile;" json:"mobile"`
	Email       string    `gorm:"type:varchar(128);column:email;unique_index;" json:"email"`
	CreateAt    time.Time `gorm:"type:datetime;column:create_at;not null;" json:"create_at"`
	LastLoginAt time.Time `gorm:"type:datetime;column:last_login_at;not null;" json:"last_login_at"`
	//UserGroups  []UserGroup
}

func (User) TableName() string {
	return "auth_user"
}

type Group struct {
	ID               int       `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Name             string    `gorm:"type:varchar(32);column:name;unique_index" json:"name"`
	CreateAt         time.Time `gorm:"type:datetime;column:create_at;" json:"create_at"`
	GroupPermissions []GroupPermission
}

func (Group) TableName() string {
	return "auth_group"
}

type Permission struct {
	ID       int       `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Action   string    `gorm:"type:varchar(32);column:action;unique_index" json:"action"`
	CreateAt time.Time `gorm:"type:datetime;column:create_at;" json:"create_at"`
}

func (Permission) TableName() string {
	return "auth_permission"
}

type UserGroup struct {
	ID      int   `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	UserID  int   `gorm:"index;column:user_id;"`
	GroupID int   `gorm:"index;column:group_id;"`
	User    User  `gorm:"foreignkey:UserID;PRELOAD:false;" json:"user"`
	Group   Group `gorm:"foreignkey:GroupID;" json:"group"`
}

func (UserGroup) TableName() string {
	return "auth_user_group_relation"
}

type GroupPermission struct {
	ID            int        `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	PerminssionID int        `gorm:"index;column:permission_id;"`
	GroupID       int        `gorm:"index;column:group_id;"`
	Permission    Permission `gorm:"foreignkey:PermissionID;" json:"permission"`
	Group         Group      `gorm:"foreignkey:GroupID;PRELOAD:false;" json:"group"`
}

func (GroupPermission) TableName() string {
	return "auth_group_perminssion_relation"
}

type KubernetesInfo struct {
	ID         int       `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	ClusterURL string    `gorm:"type:varchar(128);column:cluster_url;not null;" json:"cluster_url"`
	Port       int       `gorm:"type:integer;column:port;not null;" json:"port"`
	Username   string    `gorm:"type:varchar(64);column:username;not null;" json:"username"`
	Token      string    `gorm:"type:varchar(2048);column:token;not null;" json:"token"`
	Alias      string    `gorm:"type:varchar(128);column:alias;" json:"alias"`
	CreateAt   time.Time `gorm:"type:datetime;column:create_at;" json:"create_at"`
	UserID     int       `gorm:"index;column:user_id;" json:"user_id"`
	Namespace  string    `gorm:"type:varchar(128);column:namespace;" json:"namespace"`
}

func (KubernetesInfo) TableName() string {
	return "kubernetes_cluster"
}

type MysqlInstance struct {
	ID                       int            `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Name                     string         `gorm:"type:varchar(128);column:name;not null;" json:"name"`
	InstanceType             string         `gorm:"type:varchar(32);column:instance_type;not null;" json:"instance_type"`
	Alias                    string         `gorm:"type:varchar(128);column:alias;" json:"alias"`
	KubernetesNS             string         `gorm:"type:varchar(128);column:kubernetes_ns;not null;" json:"kubernetes_ns"`
	KubernetesSVC            string         `gorm:"type:varchar(128);column:kubernetes_svc;not null;" json:"kubernetes_svc"`
	KubernetesControllerType string         `gorm:"type:varchar(128);column:k8s_controller_type;not null;" json:"k8s_controller_type"`
	KubernetesController     string         `gorm:"type:varchar(128);column:k8s_controller;not null;" json:"k8s_controller"`
	Port                     int            `gorm:"type:integer;column:port;not null;" json:"port"`
	RootPassword             string         `gorm:"type:varchar(128);column:root_password;not null;" json:"root_password"`
	IsDelete                 bool           `gorm:"type:tinyint;column:is_delete;not null;default:0;" json:"is_delete"`
	CreateAt                 time.Time      `gorm:"type:date;column:create_at;" json:"create_at"`
	UserID                   int            `gorm:"index;column:user_id;" json:"user_id"`
	KubernetesInfoID         int            `gorm:"index;column:kubernetes_id;" json:"kubernetes_id"`
	KubernetesInfo           KubernetesInfo `gorm:"foreignkey:KubernetesInfoID;" json:"kubernetes_info"`
	User                     User           `gorm:"foreignkey:UserID;" json:"user"`
	MasterID                 int            `gorm:"column:master_id;" json:"master_id"`
	MasterInstanceID         *uint
	MasterInstance           *MysqlInstance //单表,属于某个主库
	//MasterInstance            []MysqlInstance `gorm:"foreignkey:SlaveInstanceID"`  //一个主库对应多个从库
	KubernetesPod  string          `gorm:"type:varchar(256);column:kubernetes_pod" json:"kubernetes_pod"`
	MysqlSvcLabels []MysqlSvcLabel //one mysqlinstance svc have many labels
}

func (MysqlInstance) TableName() string {
	return "mysql_instance"
}

//MySQL Lables
type MysqlSvcLabel struct {
	ID              int    `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	LabelKey        string `gorm:"type:varchar(128);column:label_key;not null;" json:"label_key"`
	LabelValue      string `gorm:"type:varchar(128);column:label_value;not null;" json:"label_value"`
	MysqlInstanceID int    `gorm:"index"` //FK,belong to MysqlInstance
}

func (MysqlSvcLabel) TableName() string {
	return "mysql_svc_label"
}

func main() {
	//连接数据库
	//dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "root:root@tcp(data:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("数据库连接失败,连接信息:%s\n", dsn)
		panic("failed to connect database")
	}

	//迁移 schema 新建数据表
	//db.AutoMigrate(&User{}, &Group{}, &Permission{}, &KubernetesInfo{}, &MysqlInstance{}, &MysqlSvcLabel{}, &UserGroup{}, &GroupPermission{})
	db.AutoMigrate(&User{}, &KubernetesInfo{}, &MysqlInstance{}, &MysqlSvcLabel{})
	fmt.Printf("数据库自动迁移完成\n")

	////收到一个主库
	//db.Create(&MysqlInstance{
	//	Name: "jinzhu",
	//	InstanceType: "master",
	//	Alias: "cmdb-mysql",
	//	KubernetesNS: "xinfracloud",
	//	KubernetesSVC: "dayu-mysql",
	//	KubernetesControllerType : "StatefulSet",
	//	KubernetesController:"dayu-mysql",
	//	Port: 3306,
	//	RootPassword: "Y2xvdWRtaW5kc19zcmU=",
	//	IsDelete: false,
	//	CreateAt: time.Now(),
	//	UserID: 1,
	//	KubernetesInfoID: 2,
	//	//SlaveInstance: []MysqlInstance{Name:"jinzhu",},
	//})
	//
	////收到一个从库
	//db.Create(&MysqlInstance{
	//	Name: "jinzhu",
	//	InstanceType: "master",
	//	Alias: "cmdb-mysql",
	//	KubernetesNS: "xinfracloud",
	//	KubernetesSVC: "dayu-mysql",
	//	KubernetesControllerType : "StatefulSet",
	//	KubernetesController:"dayu-mysql",
	//	Port: 3306,
	//	RootPassword: "Y2xvdWRtaW5kc19zcmU=",
	//	IsDelete: false,
	//	CreateAt: time.Now(),
	//	UserID: 1,
	//	KubernetesInfoID: 2,
	//	//SlaveInstance: []MysqlInstance{Name:"jinzhu",},
	//})
	//
	////查询主库
	//var mysqlInstance MysqlInstance
	//db.Where("name = ? AND kubernetes_svc=?", "jinzhu", "dayu-mysql").First(&mysqlInstance)
	//fmt.Printf("查询主库结果:\n%+v", mysqlInstance)
	//
	//
	//
	////添加从库
	//db.Create(&MysqlInstance{
	//	Name: "jinzhu-slave",
	//	InstanceType: "slave",
	//	Alias: "cmdb-mysql",
	//	KubernetesNS: "xinfracloud",
	//	KubernetesSVC: "dayu-mysql",
	//	KubernetesControllerType : "StatefulSet",
	//	KubernetesController:"dayu-mysql",
	//	Port: 3306,
	//	RootPassword: "Y2xvdWRtaW5kc19zcmU=",
	//	IsDelete: false,
	//	CreateAt: time.Now(),
	//	UserID: 1,
	//	KubernetesInfoID: 2,
	//	SlaveInstance: []MysqlInstance{mysqlInstance},
	//})

	//收到一个主库
	//	instanceMaster:=MysqlInstance{
	//		Name: "master",
	//		InstanceType: "master",
	//		Alias: "cmdb-mysql",
	//		KubernetesNS: "xinfracloud",
	//		KubernetesSVC: "dayu-mysql",
	//		KubernetesControllerType : "StatefulSet",
	//		KubernetesController:"dayu-mysql",
	//		Port: 3306,
	//		RootPassword: "Y2xvdWRtaW5kc19zcmU=",
	//		IsDelete: false,
	//		CreateAt: time.Now(),
	//		UserID: 1,
	//		KubernetesInfoID: 2,
	//	}
	//	db.Create(&instanceMaster)

	//收到一个从库
	//查询其主库
	var instanceMaster MysqlInstance
	db.Where("name = ? AND kubernetes_svc=?", "master", "dayu-mysql").First(&instanceMaster)
	fmt.Printf("查询主库结果:\n%+v", instanceMaster)

	//存储从库
	instanceSlave := MysqlInstance{
		Name:                     "slave2",
		InstanceType:             "slave",
		Alias:                    "cmdb-mysql-slave2",
		KubernetesNS:             "xinfracloud",
		KubernetesSVC:            "dayu-mysql",
		KubernetesControllerType: "StatefulSet",
		KubernetesController:     "dayu-mysql",
		Port:                     3306,
		RootPassword:             "Y2xvdWRtaW5kc19zcmU=",
		IsDelete:                 false,
		CreateAt:                 time.Now(),
		UserID:                   1,
		KubernetesInfoID:         2,
		MasterInstance:           &instanceMaster,
	}
	db.Create(&instanceSlave)

}
