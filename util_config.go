package utils4g

import (
	"github.com/plusplus1/utils4g/configuration"
	"github.com/plusplus1/utils4g/env"
)

type configUtil struct{}

var configInstance = &configUtil{}

func newConfigUtil() *configUtil {
	return configInstance
}

// SetBaseDir, set base dir
func (u *configUtil) SetBaseDir(baseDir string) {

}

// GetBaseDir, get base dir
func (u *configUtil) GetBaseDir() string {
	return env.ConfDir()
}

// ReadYaml , read yaml file
func (u *configUtil) ReadYaml(confYaml string, out interface{}) error {
	return configuration.LoadYaml(confYaml, out)
}

// ReadPath , read via config path
func (u *configUtil) ReadPath(path string, out interface{}) error {
	return configuration.LoadRelativePath(path, out)
}
