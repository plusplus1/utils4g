package utils4g

import (
	"sync"
	"time"
)

import (
	"gopkg.in/mgo.v2"
)

const (
	cstMgoDbDefaultConfigPath         = "/db/mongo/" // 配置文件路径
	cstMgoDbDefaultConfigExpire       = 3600         // 一小时自动过期
	cstMgoDbDefaultConfigMaxPoolLimit = 4096
	cstMysqlDefaultConfigPath         = "/db/mysql/" // 配置文件路径
	cstMysqlDefaultConfigExpire       = 3600         // 一小时自动过期
	cstMysqlDriverName                = "mysql"
)

var (
	defaultLayoutDate = []string{
		"2006-01-02",
		"2006-1-2",
		"2006/01/02",
		"2006/1/2",
	}

	defaultLayoutDateTime = []string{
		"2006-01-02 15:04:05",
		"2006-1-2 15:04:05",
		"2006-01-02 15:4:5",
		"2006-1-2 15:4:5",
		"2006/01/02 15:04:05",
		"2006/1/2 15:04:05",
		"2006/01/02 15:4:5",
		"2006/1/2 15:4:5",
	}
)

type (
	configUtil struct{}

	debugUtil struct {
		startupDir string
		isDebug    bool
	}

	datetimeUtil struct {
		layoutDateTime []string
		layoutDate     []string
		loc            *time.Location
	}

	strUtil struct{}

	dbMgoConf struct {
		Name       string `json:"name"`
		DbName     string `json:"db"`
		Filename   string `json:"filename"`
		UpdateTime int64  `json:"update_time"`

		ConnStr   string `yaml:"connect_string" json:"connect_string"`
		PoolLimit int    `yaml:"pool_limit" json:"pool_limit"`
	}

	dbMgoUtil struct {
		sync.Mutex

		configs        map[string]*dbMgoConf
		lastConfUpdate int64
		sessions       map[string]*mgo.Session
	}

	dbMysqlConf struct {
		Filename   string            `json:"filename"`
		UpdateTime int64             `json:"update_time"`
		Host       string            `json:"host" yaml:"host"`
		Port       int               `json:"port" yaml:"port"`
		User       string            `json:"user" yaml:"user"`
		Password   string            `json:"password" yaml:"password"`
		Db         string            `json:"db" yaml:"db"`
		Params     map[string]string `json:"params" yaml:"params"`
	}

	dbMysqlUtil struct {
		sync.Mutex
		configs map[string]*dbMysqlConf
	}

	dbUtil struct {
		Mongo *dbMgoUtil
		Mysql *dbMysqlUtil
	}
)
