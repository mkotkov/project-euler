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

func SumSquareDifference(n int) int {
	sumOfSquares := 0
	squareOfSum := 0

	for i := 1; i <= n; i++ {
		sumOfSquares += i * i
		squareOfSum += i
	}
	fmt.Println("Sum of squares:", sumOfSquares)
	fmt.Println("Square of sum before squaring:", squareOfSum)
	squareOfSum *= squareOfSum
	fmt.Println("Square of sum after squaring:", squareOfSum)

	fmt.Println("Difference (squareOfSum - sumOfSquares)...")
	return squareOfSum - sumOfSquares
}

