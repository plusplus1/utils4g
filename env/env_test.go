package env

import (
	"os"
	"path/filepath"
	"testing"
)

func TestISDebug(t *testing.T) {
	_ = os.Setenv(envKeyISDebug, "1")
	vObj.initEnv()
	if !ISDebug() {
		t.Error("Assert is debug true, but false")
	}
	_ = os.Setenv(envKeyISDebug, "")
	vObj.initEnv()
	if ISDebug() {
		t.Error("Assert is debug false, but true")
	}

}

func TestISDocker(t *testing.T) {
	_ = os.Setenv(envKeyISDocker, "hello")
	vObj.initEnv()
	if !ISDocker() {
		t.Error("Assert is docker true, but false")
	}
	_ = os.Setenv(envKeyISDocker, "")
	vObj.initEnv()
	if ISDocker() {
		t.Error("Assert is docker false, but true")
	}

}

func TestBaseDir(t *testing.T) {
	rootDir, _ := os.Getwd()
	rootDir, _ = filepath.Abs(rootDir)
	if rootDir != BaseDir() {
		t.Errorf("Assert base dir is %v, but not", rootDir)
	}
}

func TestConfDir(t *testing.T) {
	rootDir, _ := os.Getwd()

	_ = os.Setenv(envKeyISDebug, "1")
	vObj.initEnv()
	if dir := filepath.Join(rootDir, dirRelConfTest); dir != ConfDir() {
		t.Errorf("Assert conf dir is %v, but not", dir)
	} else {
		t.Logf("Assert conf dir is %v when debug=%v", dir, ISDebug())
	}

	_ = os.Setenv(envKeyISDebug, "")
	vObj.initEnv()
	if dir := filepath.Join(rootDir, dirRelConf); dir != ConfDir() {
		t.Errorf("Assert conf dir is %v, but not", dir)
	} else {
		t.Logf("Assert conf dir is %v when debug=%v", dir, ISDebug())
	}
}
