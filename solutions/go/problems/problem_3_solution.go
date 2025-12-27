// Problem 3: Largest prime factor
// The prime factors of 13195 are 5, 7, 13 and 29.
// What is the largest prime factor of the number 600851475143?  

//  Link: https://projecteuler.net/problem=3

// Solution:
package problems

import (
	"project-euler-go/base"
)

// LargestPrimeFactor finds the largest prime factor of a given number
func LargestPrimeFactor(n int) int {
	largestPrimeFactor := 1
	for i := 2; i*i <= n; i++ {
		if n%i == 0 && base.IsPrime(i) {
			largestPrimeFactor = i
		}
	}
	return largestPrimeFactor
}