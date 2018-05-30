package utils4g

import (
	"io/ioutil"
	"path/filepath"
	"sync"
)

import (
	"github.com/go-yaml/yaml"
)

var (
	flagCfgBaseDir *string = nil
	cfgSetOnce             = sync.Once{}

	configInstance = &configUtil{}
)

func newConfigUtil() *configUtil {
	return configInstance
}

// SetBaseDir, set base dir
func (u *configUtil) SetBaseDir(baseDir string) {
	if baseDir != "" {
		cfgSetOnce.Do(func() {
			var dir = baseDir
			flagCfgBaseDir = &dir
		})
	}
}

// GetBaseDir, get base dir
func (u *configUtil) GetBaseDir() string {
	if flagCfgBaseDir == nil {
		panic("You must set base directory first")
	}
	return *flagCfgBaseDir
}

// ReadYaml , read yaml file
func (u *configUtil) ReadYaml(confYaml string, out interface{}) error {

	if bytes, eRead := ioutil.ReadFile(confYaml); eRead != nil {
		return eRead
	} else {
		if eDecode := yaml.Unmarshal(bytes, out); eDecode != nil {
			return eDecode
		}
	}
	return nil
}

// ReadPath , read via config path
func (u *configUtil) ReadPath(path string, out interface{}) error {
	var confYaml = filepath.Join(u.GetBaseDir(), path+".yaml")
	return u.ReadYaml(confYaml, out)
}
