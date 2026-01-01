// Problem 3: Largest prime factor
// The prime factors of 13195 are 5, 7, 13 and 29.
// What is the largest prime factor of the number 600851475143?  

//  Link: https://projecteuler.net/problem=3

// Solution:
package problems

import (
	"project-euler-go/base" // Importing base package for prime checking function
)

// LargestPrimeFactor finds the largest prime factor of a given number
func LargestPrimeFactor(n int) int {
	largestPrimeFactor := 1 	// Initialize largest prime factor to 1
	for i := 2; i*i <= n; i++ { 	// Loop through potential factors up to the square root of n
		if n%i == 0 && base.IsPrime(i) { 	// Check if i is a factor of n and if it is prime
			largestPrimeFactor = i 	// Update largest prime factor
		}
	}
	return largestPrimeFactor // Return the largest prime factor found
}