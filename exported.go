package utils4g

import (
	"path/filepath"
)

import (
	"github.com/plusplus1/utils4g/cfg"
	"github.com/plusplus1/utils4g/datetime"
	"github.com/plusplus1/utils4g/db"
	"github.com/plusplus1/utils4g/debug"
	"github.com/plusplus1/utils4g/str"
)

var (
	// DebugUtils util
	DebugUtils = debug.Util{}

	// ConfigUtils
	ConfigUtils = cfg.Util{}

	// StrUtils
	StrUtils = str.Util{}

	DB = db.Util

	DateTimeUtils = datetime.Util
)

func init() {

	if DebugUtils.IsDebug() {
		ConfigUtils.SetBaseDir(filepath.Join(DebugUtils.GetBaseDir(), "conf_test"))
	} else {
		ConfigUtils.SetBaseDir(filepath.Join(DebugUtils.GetBaseDir(), "conf"))
	}

}
