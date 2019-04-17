package mysql

import (
	"database/sql"
)

import (
	_ "github.com/go-sql-driver/mysql"
)

func Connect(path string) *sql.DB {
	return mgr.Connect(path)
}

func GetConf(path string) (cp ConnectionParameters) {
	if c := mgr.configs.GetConf(path); c != nil {
		cp = *c
	}
	return
}

func Close(db *sql.DB) {
	if db != nil {
		var st = db.Stats()
		if st.InUse > 0 {
			return
		}
		_ = db.Close()
	}

}
