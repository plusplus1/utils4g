package utils4g

import (
	"os"
	"path/filepath"
)

var (
	debugInstance = &debugUtil{}
)

func init() {
	debugInstance.startupDir, _ = os.Getwd()

	if os.Getenv("IS_DEBUG") == "1" {
		confBaseDir := filepath.Join(debugInstance.startupDir, "conf_test")
		if st, e := os.Stat(confBaseDir); e == nil && st != nil && st.IsDir() {
			debugInstance.isDebug = true
			newConfigUtil().SetBaseDir(confBaseDir)
		}
	}

	// set config directory in production environment
	if !debugInstance.isDebug {
		confBaseDir := filepath.Join(debugInstance.startupDir, "conf")
		if st, e := os.Stat(confBaseDir); e == nil && st != nil && st.IsDir() {
			newConfigUtil().SetBaseDir(confBaseDir)
		}
	}
}

func newDebugUtil() *debugUtil {
	return debugInstance
}

func (d *debugUtil) IsDebug() bool {
	return d.isDebug
}

func (d *debugUtil) GetBaseDir() string {
	return d.startupDir
}
