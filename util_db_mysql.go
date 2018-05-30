package utils4g

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"path/filepath"
	"strings"
	"time"
)

import (
	_ "github.com/go-sql-driver/mysql"
)

var (
	dbMysqlInstance = &dbMysqlUtil{
		configs: make(map[string]*dbMysqlConf),
	}
)

func newDbMysqlUtil() *dbMysqlUtil {
	return dbMysqlInstance
}

func (mysqlUtil *dbMysqlUtil) GetConf(name string) dbMysqlConf {
	if cfg := mysqlUtil.getConf(name, false); nil != cfg {
		return *cfg
	}
	return dbMysqlConf{}
}

func (mysqlUtil *dbMysqlUtil) Connect(name string) *sql.DB {
	config := mysqlUtil.getConf(name, false)
	if config == nil {
		log.Printf("[ERROR] connect mysql fail, name=%v, error=%v", name, "no config")
		return nil
	}

	db, err := sql.Open(cstMysqlDriverName, config.buildDataSourceString())
	if err != nil {
		return nil
	}
	return db
}

func (mysqlUtil *dbMysqlUtil) Close(db *sql.DB) {
	if nil != db {
		db.Close()
	}
}

func (mysqlUtil *dbMysqlUtil) formatName(name string) string {
	if !strings.HasPrefix(name, "/") {
		name = "/" + name
	}
	return name
}

func (c *dbMysqlConf) buildDataSourceString() string {
	var s = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", c.User, c.Password, c.Host, c.Port, c.Db)
	if len(c.Params) > 0 {
		var p string
		for k, v := range c.Params {
			p += fmt.Sprintf("&%s=%s", k, v)
		}
		s += "?" + p[1:]
	}
	return s
}

func (mysqlUtil *dbMysqlUtil) reloadConf(name string, cfg *dbMysqlConf) (*dbMysqlConf, error) {
	name = mysqlUtil.formatName(name)

	var filename string
	var confUtil = newConfigUtil()

	if cfg != nil && cfg.Filename != "" {
		filename = cfg.Filename
	} else {
		filename = filepath.Join(confUtil.GetBaseDir(), cstMysqlDefaultConfigPath, name+".yaml")
		cfg = &dbMysqlConf{Filename: filename}
	}

	if err := confUtil.ReadYaml(filename, cfg); err != nil {
		return nil, err
	}
	if err := cfg.checkValid(); err != nil {
		return nil, err
	}

	cfg.UpdateTime = time.Now().Unix()
	mysqlUtil.configs[name] = cfg

	return cfg, nil
}

func (mysqlUtil *dbMysqlUtil) getConf(name string, force bool) *dbMysqlConf {
	var config *dbMysqlConf = nil
	var err error

	if !force {
		if cachedConfig, ok := mysqlUtil.configs[mysqlUtil.formatName(name)]; ok && cachedConfig != nil {
			if time.Now().Unix()-cachedConfig.UpdateTime < cstMysqlDefaultConfigExpire {
				return cachedConfig
			}
			config = cachedConfig
		}
	}

	if config, err = mysqlUtil.reloadConf(name, config); err != nil {
		log.Printf("[ERROR] reload mysql cfg %v error : %v", name, err)
		return nil
	}

	log.Printf("[INFO] reload mysql cfg %v success, %v", name, config.String())
	return config
}

func (c *dbMysqlConf) String() string {
	var bytes []byte
	bytes, _ = json.Marshal(c)
	return string(bytes)
}

func (c *dbMysqlConf) checkValid() error {
	if c.Host == "" {
		return errors.New("host empty")
	}
	if c.Port == 0 {
		c.Port = 3306
	}

	if c.User == "" || c.Password == "" {
		return errors.New("authorize information missing")
	}

	if c.Db == "" {
		return errors.New("db name missing")
	}

	return nil
}
