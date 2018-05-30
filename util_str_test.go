package utils4g

import "testing"

func TestStrUtil_NewStrUUID(t *testing.T) {

	s := newStrUtil()
	t.Log(s.NewStrUUID())

}

func TestStrUtil_Md5String(t *testing.T) {

	s := newStrUtil()
	sMd5 := s.Md5String("hello world")
	t.Log("md5=", sMd5)
	if sMd5 == "5eb63bbbe01eeed093cb22bb8f5acdc3" {
		t.Log("ok")

	} else {
		t.Error("fail")
	}
}
