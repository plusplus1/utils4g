package datetime

import (
	"testing"
)

func TestFormatLayout(t *testing.T) {
	t.Logf("foamt layout = %v", Format(NowZH()))
	t.Logf("foamt layout = %v", FormatLayout(NowZH(), DefaultDateLayout))
}

func TestParseLayout(t *testing.T) {
	var cases = []string{
		"2019-04-16 18:17:03",
		"2019/04/16 18:17:03",
		"2019/4/16 18:17:3",
		"20190416181703",
		"2019-04-16",
		"2019/04/16",
		"2019/4/16",
		"20190416",
	}

	for _, c := range cases {
		if tt, e := Parse(c); e != nil {
			t.Errorf("case %v , parsed error = %v", c, e)
		} else {
			t.Logf("case %v , parsed = %v", c, tt)
		}
	}

}
