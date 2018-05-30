package utils4g

import "testing"

func TestDebugUtil_GetBaseDir(t *testing.T) {

	d := newDebugUtil()
	t.Log(d.GetBaseDir())
}

func TestDebugUtil_IsDebug(t *testing.T) {
	d := newDebugUtil()
	t.Log(d.isDebug)
}
