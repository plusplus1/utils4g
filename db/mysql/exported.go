package mysql

import (
	"database/sql"
	"log"
)

import (
	_ "github.com/go-sql-driver/mysql"
)

const (
	driverName = "mysql"
)

type Util struct{}

func (u Util) Connect(name string) *sql.DB {
	config := cfgManager.getConf(name, false)
	if config == nil {
		log.Printf("[ERROR] connect mysql fail, name=%v, error=%v", name, "no config")
		return nil
	}

	db, err := sql.Open(driverName, config.buildDataSourceString())
	if err != nil {
		return nil
	}
	return db
}

func (u Util) Close(db *sql.DB) {
	db.Close()
}
