package utils4g

import (
	"database/sql"
)

import (
	"github.com/xgo11/configuration"
	"github.com/xgo11/env"
	"github.com/xgo11/mongo4g"
	"github.com/xgo11/mysql4g"
	"github.com/xgo11/redis4g"
	"github.com/xgo11/stdlog"
	"gopkg.in/mgo.v2"
)

var StdLog = stdlog.Std

func ISDebug() bool {
	return env.ISDebug()
}

func ISDocker() bool {
	return env.ISDocker()
}
func BaseDir() string {
	return env.BaseDir()
}
func ConfDir() string {
	return env.ConfDir()
}
func LoadAbsYamlConf(yaml string, out interface{}) error {
	return configuration.LoadYaml(yaml, out)
}

func LoadRelativePathConf(path string, out interface{}) error {
	return configuration.LoadRelativePath(path, out)
}

func ConnectMongo(path string) *mgo.Database {
	return mongo4g.Connect(path)
}

func CloseMongo(database *mgo.Database) {
	mongo4g.Close(database)
}

func ConnectMySQL(path string) *sql.DB {
	return mysql4g.Connect(path)
}
func CloseMySQL(db *sql.DB) {
	mysql4g.Close(db)
}

func ConnectRedis(path string) *redis4g.WrapClient {
	return redis4g.Connect(path)
}
