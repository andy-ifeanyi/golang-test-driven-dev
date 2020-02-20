package leapyear

import "testing"

func TestIsLeapYear(t *testing.T) {
	cases := []struct {
		in   uint16
		want bool
	}{
		{1980, true},
		{2022, false},
		{2020, true},
		{1999, false},
		{2011, false},
		{1996, true},
		{2016, true},
	}
	for _, c := range cases {
		result := IsLeapYear(c.in)
		if result != c.want {
			t.Errorf("IsLeapYear(%q) == %t, want %t", c.in, result, c.want)
		}
	}
}
