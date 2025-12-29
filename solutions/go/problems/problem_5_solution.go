// Problem 5: Smallest multiple
// Find the smallest positive number that is evenly divisible by all of the numbers from 1 to 20.

// Solution:
package problems
	
import (
	"fmt"
)
// SmallestMultiple returns the smallest positive number that is evenly divisible by all of the numbers from 1 to 20.
func SmallestMultiple(end int) int {
	start := 1
	if start < 1 || end < start {
		fmt.Println("Invalid range: numbers must be >= 1 and end >= start")
		return -1
	}

	result := 1
	for i := start; i <= end; i++ {
		result = lcm(result, i)
	}
	return result
}

// Helper function to compute least common multiple (LCM)
func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

// Helper function to compute greatest common divisor (GCD)
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
