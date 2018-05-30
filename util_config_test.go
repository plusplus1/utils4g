package utils4g

import (
	"path/filepath"
	"testing"
)

var (
	path = "conf_test"
)

func TestConfigUtil_GetBaseDir(t *testing.T) {
	cfg := newConfigUtil()
	func() {
		t.Log(cfg.GetBaseDir())
	}()

	defer func() {
		if e := recover(); e != nil {
			t.Errorf("%v", e)
		}
	}()
	cfg.SetBaseDir(path)
	t.Log(cfg.GetBaseDir(), cfg.GetBaseDir() != path)

}

func TestConfigUtil_ReadYaml(t *testing.T) {
	cfg := newConfigUtil()

	out := struct {
		ConnectString string `yaml:"connect_string"`
		PoolLimit     string `yaml:"pool_limit"`
	}{}

	yamlPath := filepath.Join(cfg.GetBaseDir(), "db/mongo/demo.yaml")
	t.Log(yamlPath)
	if e := cfg.ReadYaml(yamlPath, &out); e != nil {
		t.Error(e)
		return
	}
	t.Log(out.ConnectString)
	t.Log(out.PoolLimit)
}

func TestConfigUtil_ReadPath(t *testing.T) {
	cfg := newConfigUtil()
	out := struct {
		ConnectString string `yaml:"connect_string"`
		PoolLimit     string `yaml:"pool_limit"`
	}{}

	if e := cfg.ReadPath("db/mongo/demo", &out); e != nil {
		t.Error(e)
		return
	}
	t.Log(out.ConnectString)
	t.Log(out.PoolLimit)
}
