package datetime

import "testing"

func TestNow(t *testing.T) {
	t.Logf("now = %v", Now())
}

func TestNowZH(t *testing.T) {
	t.Logf("now zh  = %v", NowZH())
}

func TestNowUnix(t *testing.T) {
	t.Logf("now unix = %v", NowUnix())
}

func TestNowUnixNano(t *testing.T) {
	t.Logf("now unix nano = %v", NowUnixNano())
}
