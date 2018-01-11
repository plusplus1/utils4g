package debug

import (
	"testing"
)

func TestUtil_GetBaseDir(t *testing.T) {

	var u = &Util{}
	var path = u.GetBaseDir()

	t.Log(path)

}

func TestUtil_IsDebug(t *testing.T) {
	var u = &Util{}
	var isDebug bool
	isDebug = u.IsDebug()

	t.Log(isDebug)

}
