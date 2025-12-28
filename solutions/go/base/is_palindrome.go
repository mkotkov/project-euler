// IsPalindrome checks if a number is a palindrome
package base

import (
	"strconv"
)

// IsPalindrome checks if a number is a palindrome
func IsPalindrome(n int) bool {
	s := strconv.Itoa(n)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}