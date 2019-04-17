package datetime

import "testing"

func TestFromUnix(t *testing.T) {

	t.Logf("from unix = %v", FromUnix(1555409241, BeijingTZ()))

}

func TestFromUnixNano(t *testing.T) {
	t.Logf("from unix = %v", FromUnixNano(1555409241, 918400000, BeijingTZ()))
}
