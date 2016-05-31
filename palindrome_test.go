package palindrome

import (
	"fmt"
	"testing"
)

func TestPalindrome(t *testing.T) {
	cases := []struct {
		in   string
		want bool
	}{
		{"A man, a plan, a canal, Panama!", true},
		{"Hello, 世界", false},
		{"", false},
		{"1221", false},
		{"    ", false},
		{"step on no pets", true},
	}
	for _, c := range cases {
		got := isPalindrome(c.in)
		if got != c.want {
			t.Errorf("isPalindrome(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func ExamplePalindrome() {
	fmt.Println(isPalindrome("No 'x' in Nixon!"))
	// Output: true
}

func ExampleNotPalindrome() {
	fmt.Println(isPalindrome("Eeee Iiii ooo you!"))
	// Output: false
}
