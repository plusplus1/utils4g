package redis

import "testing"

func TestNewRedisConfig(t *testing.T) {

	var c = NewClient("demo")

	var key, val = "aaaa", "xxxxxx"

	c.Delete(key)

	if c.Get(key) == "" {
		t.Log("Assert key empty ok")
	} else {
		t.Errorf("key value is not empty")
	}

	c.Set(key, val)

	if c.Get(key) == val {
		t.Log("Assert key's val == val ok")
	} else {
		t.Errorf("key's val is not as expected")
	}

}
