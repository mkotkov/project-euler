// Problem 5: Smallest multiple
// Find the smallest positive number that is evenly divisible by all of the numbers from 1 to 20.

// Solution:
package problems
	
import (
	"fmt"
)
// SmallestMultiple returns the smallest positive number that is evenly divisible by all of the numbers from 1 to 20.
func SmallestMultiple(end int) int {
	start := 1 // Starting number of the range
	if start < 1 || end < start { // Validate input range
		fmt.Println("Invalid range: numbers must be >= 1 and end >= start")
		return -1 
	}

	result := 1 // Initialize result variable
	for i := start; i <= end; i++ { // Loop through numbers from start to end
		result = lcm(result, i) // Calculate LCM of result and current number
	}
	return result // Return the smallest multiple found
}

// Helper function to compute least common multiple (LCM)
func lcm(a, b int) int {
	return a * b / gcd(a, b) // LCM formula using GCD
}

// Helper function to compute greatest common divisor (GCD)
func gcd(a, b int) int {
	for b != 0 { // Euclidean algorithm
		a, b = b, a%b // Swap values
	}
	return a // Return GCD
}
