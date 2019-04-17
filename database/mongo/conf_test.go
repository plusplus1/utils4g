package mongo

import "testing"

func TestNewConnectionParams(t *testing.T) {
	if cp, err := NewConnectionParams("demo"); err != nil {
		t.Errorf("load mongo connection params fail, %v", err)
	} else {
		t.Logf("load mongo connection params ok, %v", cp)
		t.Log(cp.JSON(4))

	}
}
