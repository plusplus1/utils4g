package utils4g

import (
	"fmt"
	"testing"
)

func TestDbMgoUtil_Connect(t *testing.T) {
	m := newDbMgoUtil()
	c := m.GetConf("demo")
	fmt.Println(c.String())
	db := m.Connect("demo")
	defer m.Close(db)

	fmt.Println(db.CollectionNames())

}
