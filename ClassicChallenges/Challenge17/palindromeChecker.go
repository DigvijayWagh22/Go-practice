package main

import "strings"

func IsPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}
	s = strings.ToLower(s)
	s = strings.ReplaceAll(s, " ", "")

	var str strings.Builder
	for _, r := range s {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') {
			str.WriteRune(r)
		}
	}

	s = str.String()

	if len(s) == 0 {
		return true
	}

	for i := 0; i <= len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}

	return true

}
