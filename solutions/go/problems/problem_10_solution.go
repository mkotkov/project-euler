// Problem 10: Sum of primes below two million
// The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.
// Find the sum of all the primes below two million.

// Link: https://projecteuler.net/problem=10

// Solution:
package problems

import (
	"project-euler-go/base"
)

// SumOfPrimesBelow calculates the sum of all prime numbers below the given limit.
func SumOfPrimesBelow(limit int) int {
	sum := 0 // Initialize sum to 0
	for i := 2; i < limit; i++ { // Loop through all numbers below the limit
		if base.IsPrime(i) { // Check if the number is prime
			sum += i // Add prime number to sum
		}
	}
	return sum // Return the total sum of primes
}