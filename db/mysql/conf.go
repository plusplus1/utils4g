package mysql

import (
	"encoding/json"
	"errors"
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
	baseConfigPath   = "/db/mysql/" // 配置文件路径
	baseConfigExpire = 3600         // 一小时自动过期
)

var (
	// 初始化内部mongodb配置管理池
	cfgManager = &dbConfManager{configs: make(map[string]*dbConf)}
)

type dbConf struct {
	Filename   string            `json:"filename"`
	UpdateTime int64             `json:"update_time"`
	Host       string            `json:"host" yaml:"host"`
	Port       int               `json:"port" yaml:"port"`
	User       string            `json:"user" yaml:"user"`
	Password   string            `json:"password" yaml:"password"`
	Db         string            `json:"db" yaml:"db"`
	Params     map[string]string `json:"params" yaml:"params"`
}

// 定义db配置管理器
type dbConfManager struct {
	LastUpdateTime int64
	configs        map[string]*dbConf
}

func formatName(name string) string {
	if !strings.HasPrefix(name, "/") {
		name = "/" + name
	}
	return name
}

func (c *dbConf) String() string {
	var bytes []byte
	bytes, _ = json.Marshal(c)
	return string(bytes)
}

func (c *dbConf) checkValid() error {
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

func (c *dbConf) buildDataSourceString() string {
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

func (mgr *dbConfManager) reloadConf(name string, cfg *dbConf) (*dbConf, error) {
	name = formatName(name)
	var filename string
	var confUtil = conf.Util{}

	if cfg != nil {
		filename = cfg.Filename
	} else {
		filename = filepath.Join(confUtil.GetBaseDir(), baseConfigPath, name+".yaml")
		cfg = &dbConf{Filename: filename}
	}

	if err := confUtil.ReadYaml(filename, cfg); err != nil {
		return nil, err
	}
	if err := cfg.checkValid(); err != nil {
		return nil, err
	}

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
		log.Printf("[ERROR] reload mysql conf %v error : %v", name, err)
		return nil
	}

	log.Printf("[INFO] reload mysql conf %v success, %v", name, config.String())
	return config

}
