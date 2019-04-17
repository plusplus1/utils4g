package mysql

import "testing"

func TestNewConnectionParameters(t *testing.T) {
	if cp, err := NewConnectionParameters("demo"); err != nil {
		t.Errorf("load mysql connection fail, %v", err)
	} else {
		t.Logf("load mysql connection ok, %v", cp)
		t.Logf("connection string = %s", cp.BuildConnectionString())
		t.Log(cp.JSON(4))
	}
}
