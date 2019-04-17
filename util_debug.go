package utils4g

import (
	"github.com/plusplus1/utils4g/env"
)

type debugUtil struct{}

var debugInstance = &debugUtil{}

func newDebugUtil() *debugUtil {
	return debugInstance
}

func (d *debugUtil) IsDebug() bool {
	return env.ISDebug()
}

func (d *debugUtil) GetBaseDir() string {
	return env.BaseDir()
}
