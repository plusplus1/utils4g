package debug

import (
	"os"
	"path/filepath"
)

var (
	flagIsDebug *int    = nil
	flagPwd     *string = nil
)

func init() {
	path, _ := os.Getwd()
	flagPwd = &path

	var u = &Util{}
	u.IsDebug()
}

// Util , debug util
type Util struct{}

// IsDebug, is debug
func (d *Util) IsDebug() bool {
	for flagIsDebug == nil {
		var flag int
		flagIsDebug = &flag

		confTest := filepath.Join(*flagPwd, "conf_test")
		if info, err := os.Stat(confTest); err == nil && info != nil && info.IsDir() {
			flag = 1
			flagIsDebug = &flag
			break
		}

		if os.Getenv("IS_DEBUG") == "1" {
			flag = 1
			flagIsDebug = &flag
			break
		}
		break
	}

	return *flagIsDebug == 1
}

// getBaseDir, get execute root directory
func (d *Util) GetBaseDir() string {
	return *flagPwd
}
