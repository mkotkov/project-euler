// Problem 4: Largest palindrome product
// A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
// Find the largest palindrome made from the product of two 3-digit numbers.

//  Link: https://projecteuler.net/problem=4

// Solution:
package problems

import (
	"fmt"
	"project-euler-go/base"
)

// LargestPalindrome finds the largest palindrome made from the product of two n-digit numbers
// n_digt is length of the digits (3 for 3-digit numbers)
func LargestPalindrome(n_digt int) int {
	fmt.Printf("Writen number of digits (n): %d\n", n_digt)
	maxPalindrome := 0
	lengthLimit := 1

	for i := 0; i < n_digt; i++ {
		lengthLimit *= 10
	}

	fmt.Printf("Length limit for %d-digit numbers: %d-\n", n_digt, lengthLimit-1)
	for i := lengthLimit - 1; i >= lengthLimit/10; i-- {
		for j := lengthLimit - 1; j >= lengthLimit/10; j-- {
			product := i * j
			if base.IsPalindrome(product) && product > maxPalindrome {
				maxPalindrome = product
			}
		}
	}
	return maxPalindrome
}