package palindrome

import "strings"

// ReverseChars  ...function reverses a string (using rune).
func ReverseChars(s string) string {
	chars := []rune(s)
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}

// IsPalindrome  ...function checks if a string is a palindrome.
func IsPalindrome(s string) bool {
	if strings.ToLower(s) == ReverseChars(strings.ToLower(s)) {
		return true
	}
	return false
}
