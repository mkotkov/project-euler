//Problem 7: 10001st prime
// By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.
// What is the 10,001st prime number?

//Solution:
package problems

import (
	"project-euler-go/base"
)

// NthPrime finds the nth prime number
func NthPrime(n int) int {
	count := 0
	num := 1
	for count < n {
		num++
		if base.IsPrime(num) {
			count++
		}
	}
	return num
}

