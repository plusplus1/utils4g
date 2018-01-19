package mongo

import (
	"gopkg.in/mgo.v2-unstable"
)

type Util struct{}

func (u Util) Connect(name string) *mgo.Database {
	return sessionMgr.getDatabase(name)
}

func (u Util) Close(db *mgo.Database) {
	sessionMgr.closeSession(db)
}

func (u Util) GetConf(name string) DbConf {
	if cfg := cfgManager.getConf(name, false); cfg != nil {
		return *cfg
	} else {
		return DbConf{}
	}
}
