package utils4g

import (
	"github.com/plusplus1/utils4g/configuration"
	"github.com/plusplus1/utils4g/env"
	"github.com/plusplus1/utils4g/stdlog"
)

var StdLog = stdlog.Std

func ISDebug() bool {
	return env.ISDebug()
}

func ISDocker() bool {
	return env.ISDocker()
}
func BaseDir() string {
	return env.BaseDir()
}
func ConfDir() string {
	return env.ConfDir()
}
func LoadAbsYamlConf(yaml string, out interface{}) error {
	return configuration.LoadYaml(yaml, out)
}

func LoadRelativePathConf(path string, out interface{}) error {
	return configuration.LoadRelativePath(path, out)
}
