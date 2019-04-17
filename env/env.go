package env

import (
	"os"
	"path/filepath"
	"strconv"
	"sync"
)

const (
	envKeyISDebug  = "IS_DEBUG"
	envKeyISDocker = "KUBERNETES_PORT"
	dirRelConf     = "conf"
	dirRelConfTest = "conf_test"
)

type variables struct {
	vFlagDebug        bool
	vFlagDocker       bool
	vStrBaseDirectory string
	vStrConfDirectory string

	once sync.Once
}

var vObj = &variables{}

func (v *variables) initEnv() {
	if i, _ := strconv.Atoi(os.Getenv(envKeyISDebug)); i > 0 {
		v.vFlagDebug = true
	} else {
		v.vFlagDebug = false
	}

	if s := os.Getenv(envKeyISDocker); s != "" {
		v.vFlagDocker = true
	} else {
		v.vFlagDocker = false
	}

	v.vStrBaseDirectory, _ = os.Getwd()
	v.vStrBaseDirectory, _ = filepath.Abs(v.vStrBaseDirectory)
	var confTestDir = filepath.Join(v.vStrBaseDirectory, dirRelConfTest)
	if info, err := os.Stat(confTestDir); err == nil {
		if info.IsDir() { // if conf_test exists, set debug true
			v.vFlagDebug = true
		}
	}
	if !v.vFlagDebug {
		v.vStrConfDirectory = filepath.Join(v.vStrBaseDirectory, dirRelConf)
	} else {
		v.vStrConfDirectory = confTestDir
	}
}

func ISDebug() bool {
	vObj.once.Do(vObj.initEnv)
	return vObj.vFlagDebug
}

func ISDocker() bool {
	vObj.once.Do(vObj.initEnv)
	return vObj.vFlagDocker
}

func BaseDir() string {
	vObj.once.Do(vObj.initEnv)
	return vObj.vStrBaseDirectory
}

func ConfDir() string {
	vObj.once.Do(vObj.initEnv)
	return vObj.vStrConfDirectory
}
