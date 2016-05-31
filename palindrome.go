package palindrome

import (
	"regexp"
	"strings"
)

// https://github.com/golang/example/blob/master/stringutil/reverse.go
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func isPalindrome(p string) bool {
	// Assuming latin characters
	reg, _ := regexp.Compile("[^A-Za-z]+")
	safe := reg.ReplaceAllString(p, "")
	safe = strings.ToLower(strings.TrimSpace(safe))

	if len(safe) < 1 {
		return false
	}

	// fmt.Println(safe)
	if safe == Reverse(safe) {
		return true
	}
	return false
}
