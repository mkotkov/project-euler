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
	maxPalindrome := 0 // Initialize maximum palindrome variable
	lengthLimit := 1 // Initialize length limit variable

	for i := 0; i < n_digt; i++ {	// Calculate length limit for n-digit numbers
		lengthLimit *= 10 // 10^n_digt
	}

	fmt.Printf("Length limit for %d-digit numbers: %d-\n", n_digt, lengthLimit-1)
	for i := lengthLimit - 1; i >= lengthLimit/10; i-- { // Loop through n-digit numbers in descending order
		for j := lengthLimit - 1; j >= lengthLimit/10; j-- { // Nested loop for second n-digit number
			product := i * j // Calculate product
			if base.IsPalindrome(product) && product > maxPalindrome { // Check if product is palindrome and greater than current max
				maxPalindrome = product // Update maximum palindrome
			}
		}
	}
	return maxPalindrome // Return the largest palindrome found
}