package mongo

import (
	"encoding/json"
	"fmt"
	"log"
	"path/filepath"
	"strings"
	"time"
)
import (
	"github.com/plusplus1/utils4g/conf"
)

const (
	baseConfigPath         = "/db/mongo/" // 配置文件路径
	baseConfigExpire       = 3600         // 一小时自动过期
	baseConfigMaxPoolLimit = 4096
)

// 定义db连接配置
type dbConf struct {
	Name       string `json:"name"`
	DbName     string `json:"db"`
	Filename   string `json:"filename"`
	UpdateTime int64  `json:"update_time"`

	ConnStr   string `yaml:"connect_string" json:"connect_string"`
	PoolLimit int    `yaml:"pool_limit" json:"pool_limit"`
}

// 定义db配置管理器
type dbConfManager struct {
	LastUpdateTime int64
	configs        map[string]*dbConf
}

func (c *dbConf) String() string {
	var bytes []byte
	bytes, _ = json.Marshal(c)
	return string(bytes)
}

var (
	// 初始化内部mongodb配置管理池
	cfgManager = &dbConfManager{configs: make(map[string]*dbConf)}
)

func (mgr *dbConfManager) reloadConf(name string, cfg *dbConf) (*dbConf, error) {
	name = formatName(name)
	var filename string
	var confUtil = conf.Util{}

	if cfg != nil {
		filename = cfg.Filename
	} else {
		filename = filepath.Join(confUtil.GetBaseDir(), baseConfigPath, name+".yaml")
		cfg = &dbConf{Filename: filename, Name: name}
	}

	if err := confUtil.ReadYaml(filename, cfg); err != nil {
		return nil, err
	}

	if cfg.ConnStr == "" {
		return nil, fmt.Errorf("connect_string配置为空")
	}
	if cfg.PoolLimit > baseConfigMaxPoolLimit {
		cfg.PoolLimit = baseConfigMaxPoolLimit
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
	mgr.configs[name] = cfg

	return cfg, nil
}

func (mgr *dbConfManager) getConf(name string, force bool) *dbConf {
	var config *dbConf = nil
	var err error

	if !force {
		if cachedConfig, ok := mgr.configs[formatName(name)]; ok && cachedConfig != nil {
			if time.Now().Unix()-cachedConfig.UpdateTime < baseConfigExpire {
				return cachedConfig
			}
			config = cachedConfig
		}
	}

	if config, err = mgr.reloadConf(name, config); err != nil {
		log.Printf("[ERROR] reload mongodb conf %v error : %v", name, err)
		return nil
	}

	log.Printf("[INFO] reload mongo conf %v success, %v", name, config.String())
	return config

}
