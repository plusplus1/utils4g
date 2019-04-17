package utils4g

import (
	"github.com/plusplus1/utils4g/database/mongo"
)

import (
	"gopkg.in/mgo.v2"
)

type dbMgoUtil struct{}

var dbMgoInstance = &dbMgoUtil{}

func newDbMgoUtil() *dbMgoUtil {
	return dbMgoInstance
}

func (mgoUtil dbMgoUtil) Connect(name string) *mgo.Database {
	return mongo.Connect(name)
}

func (mgoUtil dbMgoUtil) Close(db *mgo.Database) {
	mongo.Close(db)
}
