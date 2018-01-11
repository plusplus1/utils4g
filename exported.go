package utils4g

import (
	"path/filepath"
)

import (
	"github.com/plusplus1/utils4g/conf"
	"github.com/plusplus1/utils4g/db"
	"github.com/plusplus1/utils4g/debug"
	"github.com/plusplus1/utils4g/str"
)

var (
	// DebugUtils util
	DebugUtils = debug.Util{}

	// ConfigUtils
	ConfigUtils = conf.Util{}

	// StrUtils
	StrUtils = str.Util{}

	DB = db.Util
)

func init() {

	if DebugUtils.IsDebug() {
		ConfigUtils.SetBaseDir(filepath.Join(DebugUtils.GetBaseDir(), "conf_test"))
	} else {
		ConfigUtils.SetBaseDir(filepath.Join(DebugUtils.GetBaseDir(), "conf"))
	}

}
