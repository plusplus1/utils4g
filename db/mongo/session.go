package mongo

import (
	"log"
	"sync"
)

import (
	"gopkg.in/mgo.v2-unstable"
)

// mongodb连接管理器
type sessionManager struct {
	sync.Mutex

	sessions map[string]*mgo.Session
}

var (
	// 初始化session管理器
	sessionMgr = &sessionManager{sessions: make(map[string]*mgo.Session)}
)

func (mgr *sessionManager) getDatabase(name string) *mgo.Database {
	sessionMgr.Lock()
	defer sessionMgr.Unlock()

	name = formatName(name)

	var connectError error
	var targetDb *mgo.Database

	for {
		cfg := cfgManager.getConf(name, false)
		if cfg == nil {
			log.Printf("[ERROR] get mongodb conf fail, name=%v, error=no mongodb config", name)
			break
		}

		if s, ok := mgr.sessions[name]; ok {
			sCopy := s.Clone()

			// ping检测一下，确保断线重连
			if connectError = sCopy.Ping(); connectError != nil {
				log.Printf("[WARN] ping mongodb and then retry, name=%v, error=%v", name, connectError)
				sCopy.Close()
				s.Close()
			} else {
				targetDb = sCopy.DB(cfg.DbName)
				break
			}
		}

		// 尝试建立新的连接
		var s *mgo.Session
		if s, connectError = mgo.Dial(cfg.ConnStr); connectError != nil {
			log.Printf("[ERROR] connect mongodb fail, name=%v, error=%v", name, connectError)
			break
		}
		if connectError = s.Ping(); connectError != nil {
			log.Printf("[ERROR] ping mongodb fail, name=%v, error=%v", name, connectError)
			break
		}

		// 连接成功
		mgr.sessions[name] = s

		sCopy := s.Clone()
		targetDb = sCopy.DB(cfg.DbName)

		break
	}

	return targetDb
}

func (mgr *sessionManager) closeSession(db *mgo.Database) {
	if db != nil {
		db.Session.Close()
	}
}

func (mgr *sessionManager) closeAllSession() {
	mgr.Lock()
	defer mgr.Unlock()

	for _, s := range mgr.sessions {
		if s != nil {
			s.Close()
		}
	}
	mgr.sessions = make(map[string]*mgo.Session)
}
