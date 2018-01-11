package cfg

import (
	"io/ioutil"
	"path/filepath"
)

import (
	"github.com/go-yaml/yaml"
)

// ReadYaml , read yaml file
func (u *Util) ReadYaml(confYaml string, out interface{}) error {

	if bytes, eRead := ioutil.ReadFile(confYaml); eRead != nil {
		return eRead
	} else {
		if eDecode := yaml.Unmarshal(bytes, out); eDecode != nil {
			return eDecode
		}
	}
	return nil
}

// ReadPath , read via config path
func (u *Util) ReadPath(path string, out interface{}) error {
	var confYaml = filepath.Join(u.GetBaseDir(), path+".yaml")
	return u.ReadYaml(confYaml, out)
}
