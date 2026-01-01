// Problem 1: Multiples of 3 and 5
// If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. 
// The sum of these multiples is 23.
// Find the sum of all the multiples of 3 or 5 below 1000.

// Link: https://projecteuler.net/problem=1


// Solution
package problems

// MultiplesOf3and5 calculates the sum of all multiples of 3 or 5 below the given number.
func MultiplesOf3and5(number int) int {
	sum := 0 // Initialize sum to 0
	for i := 0; i < number; i++ { // Loop through numbers below the given number
		if i%3 == 0 || i%5 == 0 { // Check if the number is a multiple of 3 or 5
			sum += i // Add the multiple to the sum
		}
	} 
	return sum 	// Return the final sum
}
