package main

import (
	//	"os"
	//	"path"
	"time"

	"fmt"
	"os"
	"os/exec"
	"path"
	"strconv"

	"context"

	"log"

	"github.com/arthurkiller/rollingwriter"

	// import mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
)

//export GO111MODULE=off ; go run main.go

const (
	// DbPort mysql server's port
	DbPort = "3306"
	// DbName mysql db name
	DbName = "qfa"
	// LogDir log directory
	LogDir = "/var/log"
	// DefaultUpgradePath upgrade path for save sql file
	DefaultUpgradePath = "/etc/neonsan/"
	// ShardSize default 64G
	ShardSize = (64 << 30)
	//CurrentDbVersion = 2	//neonsan v1.1
	//CurrentDbVersion = 3	//neonsan v1.2
	//CurrentDbVersion = 4	//neonsan v1.2.1, i.e. v2.0.3
	//CurrentDbVersion = 5	//neonsan v2.1
	//CurrentDbVersion = 6	//neonsan v2.1.3
	//CurrentDbVersion = 7	//neonsan v2.1.17+
	//CurrentDbVersion = 8	//neonsan v2.2.0
	//CurrentDbVersion = 9	//neonsan v2.4.0
	//CurrentDbVersion = 10	//neonsan v2.5
	//CurrentDbVersion = 11	//neonsan v3.0
	//CurrentDbVersion = 12	//neonsan v3.0.1
	//CurrentDbVersion = 13	//neonsan v3.0.2
	//CurrentDbVersion = 14 //neonsan v3.1
	//CurrentDbVersion db version
	//CurrentDbVersion = 15 //neonsan v3.1.3
	//CurrentDbVersion = 16 //neonsan v3.2
	//CurrentDbVersion = 17 //neonsan v3.2.2
	//CurrentDbVersion = 18 //neonsan v3.2.3
	//CurrentDbVersion = 19 //neonsan v3.3.0
	CurrentDbVersion = 20 //neonsan v3.3.2

	//V3DbVersion db version with neonsan 3.0, must not modify
	V3DbVersion = 11 //Warning!!! Must not modify the V3DbVersion Value

	//PING_TIMEOUT ping timeout
	PING_TIMEOUT = 10 * time.Second
	//CONNECT_TIMEOUT connect timeout
	CONNECT_TIMEOUT = "10s"
)

var engine *xorm.Engine
var dbConnStr string
var dbIp string

// EnableSqlLog log direct to qfcenter-sql.log, set level info if the value is true, set level warn if false
var EnableSqlLog = false

// Volume db volume table columns
type Volume struct {
	Id              uint64 `xorm:"pk"`
	Name            string
	Size            uint64
	RepCount        uint64
	NextSnapSeq     int32
	Status          string
	MinRepCount     uint64
	CreatedTime     time.Time
	StatusTime      time.Time
	MaxBlocksize    uint64
	PoolId          uint64
	FullName        string
	MetaVer         uint64
	Role            string
	ProvisionType   string
	Policy          string
	EnablePerfStats uint64
	stores          []uint64 `xorm:"-"` //stores cary this volume, not mapping to db field
}

// Shard db shard table columns
type Shard struct {
	Id              uint64
	VolumeId        uint64
	ShardIndex      uint64
	PrimaryRepIndex uint64
	Status          string
	Bmpalldirty     uint64
	Replicas        []*Replica `xorm:"-"`
}

// Replica db replica table columns
type Replica struct {
	Id           uint64
	VolumeId     uint64
	ShardId      uint64
	ReplicaIndex uint64
	StoreId      uint64
	SsdUuid      string
	Status       string
	StatusTime   time.Time
}

// MetroReplica db metro_replica table columns
type MetroReplica struct {
	Id         uint64
	VolumeId   uint64
	Ip         string
	Port       int
	Rpo        uint64
	Status     string
	StatusTime time.Time
}

// StretchReplica db stretch_replica table columns
type StretchReplica struct {
	Id                   uint64
	VolumeId             uint64
	VolumeName           string
	OrignalPrimaryDcName string
	CurrentPrimaryDcName string
	OrignalRepCount      int32
	SyncStatus           string
	SyncStatusTime       time.Time
}

// Store db store table columns
type Store struct {
	Id            uint64
	Name          string
	MngtIp        string
	Chassis       string
	FailureDomain string
	Position      int32
	RgId          int32
	ObjectSize    uint64
	Status        string
}

func (s *Store) String() string {
	return fmt.Sprintf("%d", s.Id)
}

// Ssd db ssd table columns
type Ssd struct {
	DevName     string
	Uuid        string
	StoreId     uint64
	Position    int32
	Capacity    uint64
	Status      string
	Free        uint64
	Version     string
	Type        string
	CreatedTime time.Time
	StatusTime  time.Time
}

// UserKey db user_key table columns
type UserKey struct {
	Id               uint64
	KeyName          string
	UserKey          string
	UserKeyEncrypted string
	UsedCount        uint64
	CreatedTime      time.Time
}

// VolumeKey db volume_key table columns
type VolumeKey struct {
	VolumeId  uint64
	Iv        string
	MasterKey string
	KeyName   string
	Encrypte  string
}

// ResourceGroup db resource_group table columns
type ResourceGroup struct {
	Id           uint64
	RgName       string
	Labels       string
	CreatedTime  time.Time
	AutoBalance  int
	MonitorIp    string
	PrometheusIp string
}

// RgNode db rg_node table columns
type RgNode struct {
	RgName  string
	RgId    uint64
	StoreId uint64
}

// Snapshot db snapshot table columns
type Snapshot struct {
	Id       uint64
	VolumeId uint64
	SnapSeq  int32
	Name     string
	Status   string
	Created  time.Time
	BlockSs  []uint8
	Size     uint64
}

// Port db port table columns
type Port struct {
	StoreId  uint64
	Position int32
	Ip       string
	Type     int32
	Status   string
}

// ErrorRecord db error_record table columns
type ErrorRecord struct {
	Id        uint64
	ErrorTime time.Time
	ErrorType string
	FromIp    string
	Message   string
}

// Pool db pool table columns
type Pool struct {
	Id          uint64
	Name        string
	CreatedTime time.Time
}

// OperationRecord db operation_record table columns
type OperationRecord struct {
	Id          uint64
	FromIp      string
	OpType      string
	ReceiveTime time.Time
	Detail      string
}

// Parameter db parameter table columns
type Parameter struct {
	Name  string
	Value string
}

// Alert db alert table columns
type Alert struct {
	Id      uint64
	Type    string
	Message string
	Created time.Time
}

// ReplicaUsage db replica_usage table columns
type ReplicaUsage struct {
	Id        uint64
	Allocated uint64
}

// Acl db acl table columns
type Acl struct {
	Id         uint64
	Priority   int32
	Host       string
	Action     string
	Volume     string
	VolumeId   uint64
	Status     string
	CreateTime time.Time
}

// FrontPortal db front_portal table columns
type FrontPortal struct {
	Id     uint64
	Portal string
	Status string
}

// VipPool db vip_pool table columns
type VipPool struct {
	ID         uint64
	VIP        string
	RgId       uint64
	Available  bool
	CreateTime time.Time
	StatusTime time.Time
}

// VipPortal db vip_portal table columns
type VipPortal struct {
	Id       uint64
	Vip      string
	EthDev   string
	Portal   string
	Type     string
	RouteId  uint64
	Priority uint64
	State    string
	Role     string
	Action   string
	DenyTime time.Time
	Status   string
}

// VolumeIfront db volume_ifront table columns
type VolumeIfront struct {
	Id             uint64
	Volumefullname string
	DevName        string
	TgtIqn         string
	IniIqn         string
	TgtPortal      string
	Vip            string
	TgtLunid       uint64
	IniLunid       uint64
	Status         string
	Tha            uint64
	Volumeuuid     string
}

// VolumeInitiator db volume_initiator table columns
type VolumeInitiator struct {
	Id             uint64
	Initiatorwwn   string
	InitiatorIp    string
	Volumefullname string
	IniLunid       uint64
	Status         string
}

// VolumeIniInfo db volume_ini_info table columns
type VolumeIniInfo struct {
	Id             uint64
	Targetwwn      string
	Initiatorwwn   string
	InitiatorIp    string
	Volumefullname string
	IniLunid       uint64
	Vip            string
	Status         string
}

// Initiator db initiator table columns
type Initiator struct {
	Id        uint64
	IniIqn    string
	MgmtIp    string
	MgmtPort  uint64
	NextLunid uint64
}

// Target db target table columns
type Target struct {
	Id         uint64
	TgtIqn     string
	Portal     string
	Vip        string
	TargetAuth string
	Owner      string
	NextLunid  uint64
}

// InitiatorTarget db initiator_target table columns
type InitiatorTarget struct {
	TgtIqn       string
	IniIqn       string
	TargetPortal string
	IniMgmtIp    string
	Tha          uint64
	Owntype      string
}

// NeonsanVersion db neonsan_version table columns
type NeonsanVersion struct {
	MngtIp       string
	BuildVersion string
	CreatedTime  time.Time
}

// CloneVolume db clone_volume table columns
type CloneVolume struct {
	Id           uint64
	SourcevolId  uint64
	TargetvolId  uint64
	Status       string
	CreatedTime  time.Time
	StatusTime   time.Time
	JobId        string
	SnapshotName string
}

// MutexGroup db mutex_group table columns
type MutexGroup struct {
	Id               uint64
	GroupName        string
	RgName           string
	MutexVolumeCount uint64
	CreatedTime      time.Time
}

// MutexVolume db mutex_volume table columns
type MutexVolume struct {
	MutexGroupName string
	VolumeId       uint64
	RgName         string
}

// CandidateDisk db candidate_disk table columns
type CandidateDisk struct {
	Id           uint64
	DevName      string
	StoreId      uint64
	Capacity     uint64
	Type         string
	SerialNumber string
	Position     int32
	BeingUsed    bool
	WearDegree   string
}

// Qos db qos table columns
type Qos struct {
	VolumeId  uint64
	Iops      uint64
	Bps       uint64
	BurstIops uint64
	BurstBps  uint64
}

// Center db center table columns
type Center struct {
	MngtIp     string
	Role       string
	ExpireTime time.Time
	RoleTime   time.Time
}

// VolumeDeleted db volume_deleted table columns
type VolumeDeleted struct {
	Id          uint64 `xorm:"pk"`
	Name        string
	DeletedTime time.Time
	PoolId      uint64
	FullName    string
}

// License db license table columns
type License struct {
	Gid  string
	Code string
}

// MysqlInit init mysql and setup engine
func MysqlInit(user string, pass string, ip string, dbName string) error {
	var err error

	// Set sql engine
	dbArgs := user + ":" + pass + "@tcp(" + ip + ":" + DbPort + ")/" + dbName + "?timeout=" + CONNECT_TIMEOUT
	dbConnStr = dbArgs
	dbIp = ip
	err = setupEngine(dbArgs, LogDir)
	if err != nil {
		log.Printf("DB set engine failed:%s. resetup engine ... ", err.Error())
		engine.Close()
		var e error
		for i := 0; i < 10; i++ {
			time.Sleep(2 * time.Second)
			if e = setupEngine(dbConnStr, LogDir); e == nil {
				break
			}
		}
		if e != nil {
			log.Fatal("DB setup failed after try 10 times")
		}
	}
	return nil
}

func setupEngine(dbArgs string, logDir string) error {
	var _error error
	if engine, _error = xorm.NewEngine("mysql", dbArgs); _error != nil {
		log.Printf("Failed to connect to database: %s", _error.Error())
	} else {
		engine.SetMapper(core.GonicMapper{})
		if EnableSqlLog {
			logPath := path.Join(logDir, "qfcenter-sql.log")
			os.MkdirAll(path.Dir(logPath), os.ModePerm)
			config := rollingwriter.Config{
				LogPath:           logDir,                      //日志路径
				TimeTagFormat:     "20060102150405",            //时间格式串
				FileName:          "qfcenter-sql",              //日志文件名
				MaxRemain:         10,                          //配置日志最大存留数
				RollingPolicy:     rollingwriter.VolumeRolling, //配置滚动策略 norolling timerolling volumerolling
				RollingVolumeSize: "50M",                       //配置截断文件下限大小
				WriterMode:        "none",
				// Compress will compress log file with gzip
				Compress: false}
			writer, err := rollingwriter.NewWriterFromConfig(&config)
			if err != nil {
				log.Printf("Failed to create xorm.log: %s", err.Error())
				return err
			}
			log.Printf("XORM log to:%s", logPath)
			l := xorm.NewSimpleLogger(writer)
			l.SetLevel(core.LOG_INFO)
			engine.SetLogger(l)
			engine.ShowSQL(true)
		} else {
			engine.Logger().SetLevel(core.LOG_WARNING)
		}

		//var err error
		var err error
		if location, err := time.LoadLocation("Asia/Shanghai"); err == nil {
			engine.TZLocation = location
		}
		engine.SetMaxIdleConns(2000)
		engine.SetMaxOpenConns(2000)
		ctx, cancel := context.WithTimeout(context.Background(), PING_TIMEOUT)
		defer cancel()
		if err := engine.PingContext(ctx); err != nil {
			log.Printf("Engine failed to ping: %s", err.Error())
			return err
		}
		log.Printf("Successed connect DB")
		// vipPortals := make([]VipPortal, 0, 9)
		// if err := engine.Where("vip!=''").Find(&vipPortals); err != nil {
		// 	log.Printf("Failed to find vip_portals : %s", err.Error())
		// 	time.Sleep(time.Second * 3)
		// }
		// log.Printf("vipPortals:%v", vipPortals)
		return err
	}
	return _error
}

// GetEngine get sql engine
func GetEngine() *xorm.Engine {
	// ctx, cancel := context.WithTimeout(context.Background(), PING_TIMEOUT)
	// defer cancel()
	// log.Printf("GetEngine")
	// if err := engine.PingContext(ctx); err != nil {
	// 	log.Printf("DB ping failed:%s. resetup engine ... ", err.Error())
	// 	engine.Close()
	// 	var e error
	// 	for i := 0; i < 3; i++ {
	// 		if i != 0 {
	// 			time.Sleep(2 * time.Second)
	// 		}
	// 		if e = setupEngine(dbConnStr, LogDir); e == nil {
	// 			break
	// 		}
	// 	}
	// 	if e != nil {
	// 		log.Fatal("DB setup failed after try 3 times")
	// 	}
	// }
	return engine
}

// GetEngineWithNoRetry get sql engine
func GetEngineWithNoRetry() *xorm.Engine {
	return engine
}

// GetParameterAsInt get parameter from engine and convert to int
func GetParameterAsInt(dbEng *xorm.Engine, name string, defaultValue int) (int, error) {
	if dbEng == nil {
		dbEng = GetEngine()
	}
	param := new(Parameter)
	exist, err := dbEng.IsTableExist(param)
	if err != nil {
		return 0, err
	}
	if !exist {
		return 0, nil
	}
	has, err := dbEng.Where("name=?", name).Get(param)
	if err != nil {
		return 0, err
	}
	if !has {
		return defaultValue, nil
	}
	value, err := strconv.ParseInt(param.Value, 10, 0)
	if err != nil {
		return 0, err
	}
	return int(value), nil
}

// GetParameterAsString get parameter from engine and convert to string
func GetParameterAsString(name string, defaultValue string) (string, error) {
	param := new(Parameter)
	_, err := GetEngine().IsTableEmpty(param)
	if err != nil {
		return "", nil
	}
	has, err := GetEngine().Where("name=?", name).Get(param)
	if err != nil {
		return "", err
	}
	if !has {
		return defaultValue, nil
	}
	return param.Value, nil
}

// UpgradeDb upgrade db command
func UpgradeDb(mysqlIp string, mysqlUser string, mysqlPassword string, dbName string, path string) error {
	if _, err := exec.Command("/bin/bash", "-c",
		fmt.Sprintf("/usr/bin/mysql -h%s -u%s -p%s %s -e 'source %s;'", mysqlIp, mysqlUser, mysqlPassword, dbName, path)).Output(); err != nil {
		fmt.Printf("Failed to exel upgrade script: %s, %s", path, err.Error())
		return err
	}
	return nil
}

// GetStores get store ids which the volume load
func (v *Volume) GetStores() ([]uint64, error) {
	if len(v.stores) > 0 {
		return v.stores, nil
	}

	ids := make([]struct{ StoreId uint64 }, 0, 64)
	if err := GetEngine().SQL("select distinct store_id from replica where volume_id=?", v.Id).Find(&ids); err != nil {
		return nil, err
	}
	v.stores = make([]uint64, len(ids), len(ids))
	for i, id := range ids {
		v.stores[i] = id.StoreId
	}
	return v.stores, nil
}

// GetShardCount get shard count of the volume
func (v *Volume) GetShardCount() int {
	return int((v.Size + ShardSize - 1) / ShardSize)
}

// GetStore get store row through the replica index
func (r *Replica) GetStore() (*Store, error) {
	var s Store
	has, err := GetEngine().Where("id=?", r.StoreId).Get(&s)
	if err != nil {
		return nil, err
	}
	if !has {
		log.Printf("store not found in db for replica:%#x", r.Id)
		return nil, nil
	}
	return &s, nil
}

// replicaToVolumeId calc volume id through replica
func replicaToVolumeId(repId uint64) uint64 {
	return repId &^ 0x0ffffff
}

// GetVolume get volume rows through the volume index
func (r *Replica) GetVolume() (*Volume, error) {
	var v Volume
	has, err := GetEngine().Where("id=?", replicaToVolumeId(r.Id)).Get(&v)
	if err != nil {
		return nil, err
	}
	if !has {
		log.Printf("volume not found in db for replica:%#x", r.Id)
		return nil, nil
	}
	return &v, nil
}

// GetReplicaById get replica row through the replica index
func GetReplicaById(repId uint64) (*Replica, error) {
	var rep Replica
	has, err := GetEngine().Where("id=?", repId).Get(&rep)
	if err != nil {
		return nil, err
	}
	if !has {
		log.Printf("replica not found in db rep_id:%#x", repId)
		return nil, nil
	}
	return &rep, nil
}

// GetReplicas get replica rows through the shard index
func (s *Shard) GetReplicas() ([]*Replica, error) {
	if s.Replicas != nil {
		return s.Replicas, nil
	}
	s.Replicas = make([]*Replica, 0, 3)
	if err := GetEngine().Where("shard_id=?", s.Id).Find(&s.Replicas); err != nil {
		return nil, err
	}
	return s.Replicas, nil
}

func main() {
	mysqlUser := "neonsan"
	mysqlPass := "zhu88jie"
	mysqlIp := "172.31.32.30"
	mysqldbName := "neonsan"
	if err := MysqlInit(mysqlUser, mysqlPass, mysqlIp, mysqldbName); err != nil {
		log.Printf("err:%s", err.Error())
	}
	vipPortals := make([]VipPortal, 0, 9)
	if err := GetEngine().Where("vip!=''").Find(&vipPortals); err != nil {
		log.Printf("Failed to find vip_portals : %s", err.Error())
		time.Sleep(time.Second * 3)
	}
	for _, vipPortal := range vipPortals {
		log.Printf("vipPortal:%s", vipPortal)
	}
	time.Sleep(1000000000 * time.Second)
}
