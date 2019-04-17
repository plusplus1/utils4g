package utils4g

import "testing"

func TestDbMgoUtil_Connect(t *testing.T) {
	util := newDbMgoUtil()
	db := util.Connect("demo")
	defer util.Close(db)

	t.Log(db.CollectionNames())

}
