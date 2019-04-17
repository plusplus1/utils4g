package utils4g

import (
	"database/sql"
)

import (
	"github.com/plusplus1/utils4g/database/mysql"
)

type dbMysqlUtil struct {
}

var dbMysqlInstance = &dbMysqlUtil{}

func newDbMysqlUtil() *dbMysqlUtil {
	return dbMysqlInstance
}

func (mysqlUtil *dbMysqlUtil) Connect(name string) *sql.DB {
	return mysql.Connect(name)
}

func (mysqlUtil *dbMysqlUtil) Close(db *sql.DB) {
	mysql.Close(db)
}
