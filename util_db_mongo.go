package utils4g

import (
	"encoding/json"
	"fmt"
	"log"
	"path/filepath"
	"strings"
	"time"
)

import (
	"gopkg.in/mgo.v2"
)

var (
	dbMgoInstance = &dbMgoUtil{
		configs:  make(map[string]*dbMgoConf),
		sessions: make(map[string]*mgo.Session),
	}
)

func newDbMgoUtil() *dbMgoUtil {
	return dbMgoInstance
}

func (c *dbMgoConf) String() string {
	var bytes []byte
	bytes, _ = json.Marshal(c)
	return string(bytes)
}

func (mgoUtil dbMgoUtil) Connect(name string) *mgo.Database {
	return mgoUtil.getDatabase(name)
}

func (mgoUtil dbMgoUtil) Close(db *mgo.Database) {
	mgoUtil.closeSession(db)
}

func (mgoUtil dbMgoUtil) GetConf(name string) dbMgoConf {
	if cfg := mgoUtil.getConf(name, false); cfg != nil {
		return *cfg
	} else {
		return dbMgoConf{}
	}
}

func (mgoUtil *dbMgoUtil) reloadConf(name string, cfg *dbMgoConf) (*dbMgoConf, error) {
	name = mgoUtil.formatName(name)

	var filename string
	var confUtil = newConfigUtil()

	if cfg != nil && cfg.Filename != "" {
		filename = cfg.Filename
	} else {
		filename = filepath.Join(confUtil.GetBaseDir(), cstMgoDbDefaultConfigPath, name+".yaml")
		cfg = &dbMgoConf{Filename: filename, Name: name}
	}

	if err := confUtil.ReadYaml(filename, cfg); err != nil {
		return nil, err
	}

	if cfg.ConnStr == "" {
		return nil, fmt.Errorf("connect_string配置为空")
	}
	if cfg.PoolLimit > cstMgoDbDefaultConfigMaxPoolLimit {
		cfg.PoolLimit = cstMgoDbDefaultConfigMaxPoolLimit
	}

	var dbName, tmpStr = "", cfg.ConnStr
	if p := strings.Index(cfg.ConnStr, "?"); p > 0 {
		tmpStr = cfg.ConnStr[0:p]
	}
	if parts := strings.Split(tmpStr, "/"); len(parts) > 1 {
		dbName = parts[len(parts)-1]
	}
	if len(dbName) < 1 {
		return nil, fmt.Errorf("格式错误，解析不到db名称")
	}
	cfg.DbName = dbName
	cfg.UpdateTime = time.Now().Unix()
	mgoUtil.configs[name] = cfg

	return cfg, nil
}

func (mgoUtil *dbMgoUtil) getConf(name string, force bool) *dbMgoConf {
	var config *dbMgoConf = nil
	var err error

	if !force {
		if cachedConfig, ok := mgoUtil.configs[mgoUtil.formatName(name)]; ok && cachedConfig != nil {
			if time.Now().Unix()-cachedConfig.UpdateTime < cstMgoDbDefaultConfigExpire {
				return cachedConfig
			}
			config = cachedConfig
		}
	}

	if config, err = mgoUtil.reloadConf(name, config); err != nil {
		log.Printf("[ERROR] reload mongodb cfg %v error : %v", name, err)
		return nil
	}

	log.Printf("[INFO] reload mongo cfg %v success, %v", name, config.String())
	return config

}

func (mgoUtil *dbMgoUtil) formatName(name string) string {
	if !strings.HasPrefix(name, "/") {
		name = "/" + name
	}
	return name
}

func (mgoUtil *dbMgoUtil) getDatabase(name string) *mgo.Database {
	mgoUtil.Lock()
	defer mgoUtil.Unlock()

	name = mgoUtil.formatName(name)

	var connectError error
	var targetDb *mgo.Database

	for {
		cfg := mgoUtil.getConf(name, false)
		if cfg == nil {
			log.Printf("[ERROR] get mongodb cfg fail, name=%v, error=no mongodb config", name)
			break
		}

		if s, ok := mgoUtil.sessions[name]; ok {
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
		mgoUtil.sessions[name] = s

		sCopy := s.Clone()
		targetDb = sCopy.DB(cfg.DbName)

		break
	}

	return targetDb
}

func (mgoUtil *dbMgoUtil) closeSession(db *mgo.Database) {
	if db != nil {
		db.Session.Close()
	}
}

func (mgoUtil *dbMgoUtil) closeAllSession() {
	mgoUtil.Lock()
	defer mgoUtil.Unlock()

	for _, s := range mgoUtil.sessions {
		if s != nil {
			s.Close()
		}
	}
	mgoUtil.sessions = make(map[string]*mgo.Session)
}
