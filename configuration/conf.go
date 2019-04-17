package configuration

import (
	"io/ioutil"
	"path"
	"path/filepath"
)

import (
	"github.com/go-yaml/yaml"
	"github.com/plusplus1/utils4g/env"
)

func LoadYaml(absoluteFilepath string, out interface{}) (err error) {
	var bs []byte
	if bs, err = ioutil.ReadFile(absoluteFilepath); err == nil {
		err = yaml.Unmarshal(bs, out)
	}
	return
}

func LoadRelativePath(relativeYamlPath string, out interface{}) error {
	if ext := path.Ext(relativeYamlPath); ext == "" {
		return LoadYaml(filepath.Join(env.ConfDir(), relativeYamlPath+".yaml"), out)
	} else {
		return LoadYaml(filepath.Join(env.ConfDir(), relativeYamlPath), out)
	}
}
