package utils4g

import (
	"fmt"
	"testing"

	mysqlImpl "github.com/go-sql-driver/mysql"
)

func TestDbMysqlUtil_Connect(t *testing.T) {

	m := newDbMysqlUtil()

	db := m.Connect("demo")
	defer m.Close(db)

	rs, err := db.Query("show tables;")
	if err != nil {
		var i interface{} = err
		if e, ok := i.(*mysqlImpl.MySQLError); ok {
			t.Log(e.Number)
			t.Log(e.Message)
		} else {
			fmt.Println()
		}
		return
	}

	defer rs.Close()

	fmt.Println(rs.Columns())

	var tableName string
	for rs.Next() {
		if e := rs.Scan(&tableName); e != nil {
			t.Error(e)
			break
		}
		fmt.Println(tableName)
	}

}
