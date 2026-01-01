// Problem 6: Sum square difference
// The sum of the squares of the first ten natural numbers is,
// 1^2 + 2^2 + ... + 10^2 = 385
// The square of the sum of the first ten natural numbers is,
// (1 + 2 + ... + 10)^2 = 55^2 = 3025
// Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is,
// 3025 - 385 = 2640
// Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.

// Solution:
package problems

import "fmt"

// SumSquareDifference returns the difference between the sum of the squares of the first n natural numbers and the square of the sum.
func SumSquareDifference(n int) int {
	sumOfSquares := 0 // Initialize sum of squares variable
	squareOfSum := 0 // Initialize square of sum variable

 	for i := 1; i <= n; i++ { // Loop through first n natural numbers
		sumOfSquares += i * i // Calculate sum of squares
		squareOfSum += i // Calculate sum for square of sum
	}
	fmt.Println("Sum of squares:", sumOfSquares)
	fmt.Println("Square of sum before squaring:", squareOfSum)
	squareOfSum *= squareOfSum // Square the sum
	fmt.Println("Square of sum after squaring:", squareOfSum)

	fmt.Println("Difference (squareOfSum - sumOfSquares)...")
	return squareOfSum - sumOfSquares // Return the difference
}

