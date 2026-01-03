// Problem 9: Special Pythagorean triplet

// A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,
// a^2 + b^2 = c^2
// For example, 3^2 + 4^2 = 9 + 16 = 25 = 5^2.
// There exists exactly one Pythagorean triplet for which a + b + c = 1000.
// Find the product abc.

// Link: https://projecteuler.net/problem=9

// Solution:
package problems
import "fmt"

// SpecialPythagoreanTriplet finds the product abc of the Pythagorean triplet for which a + b + c = sumTotal.
func SpecialPythagoreanTriplet(sumTotal int) int {
	for a := 1; a < sumTotal/3; a++ { // Loop through possible values of a
		for b := a + 1; b < sumTotal/2; b++ { // Loop through possible values of b
			c := sumTotal - a - b // Calculate c based on the sum total
			if a*a + b*b == c*c { // Check if a, b, c form a Pythagorean triplet
				fmt.Printf("Found triplet: a=%d, b=%d, c=%d\n", a, b, c)
				return a * b * c // Return the product abc
			}
		}
	}
	fmt.Println("No triplet found")
	return -1 // Return -1 if no triplet is found
}

