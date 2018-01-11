package db

import (
	"github.com/plusplus1/utils4g/db/mongo"
	"github.com/plusplus1/utils4g/db/mysql"
)

var (
	Util = util{
		MySQL: mysql.Util{},
		Mongo: mongo.Util{},
	}
)

type util struct {
	MySQL mysql.Util
	Mongo mongo.Util
}
