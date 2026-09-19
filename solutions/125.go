package solutions

import "strings"

func isPalindrome(s string) bool {
	s = strings.ToLower(s)

	var s1 strings.Builder
	for _, v := range s {
		if (v >= 'a' && v <= 'z') || (v >= '0' && v <= '9') {
			s1.WriteRune(v)
		}
	}
	s = s1.String()

	for i, j := 0, len(s)-1; i < len(s)/2; {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}

	return true
}
