package utils4g

import (
	"fmt"
	"testing"
)

func TestDbMysqlUtil_Connect(t *testing.T) {

	m := newDbMysqlUtil()
	c := m.GetConf("demo")
	fmt.Println(c.String())

	db := m.Connect("demo")
	defer m.Close(db)

	rs, _ := db.Query("show tables;")
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
