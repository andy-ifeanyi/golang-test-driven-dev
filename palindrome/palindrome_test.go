package palindrome

import "testing"

func TestIsPalindrome(t *testing.T) {
	cases := []struct {
		in   string
		want bool
	}{
		{"ciVic", true},
		{"true", false},
		{"madam", true},
		{"renegade", false},
		{"Guava", false},
		{"tattat", true},
		{"02022020", true},
	}
	for _, c := range cases {
		result := IsPalindrome(c.in)
		if result != c.want {
			t.Errorf("IsPalindrome(%q) == %t, want %t", c.in, result, c.want)
		}
	}
}
