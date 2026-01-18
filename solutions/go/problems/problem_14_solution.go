//Problem 14: Longest Collatz Sequence
//The following iterative sequence is defined for the set of positive integers:
//
//n → n/2 (n is even)
//n → 3n + 1 (n is odd)
//
//Using the rule above and starting with 13, we generate the following sequence:
//
//13 → 40 → 20 → 10 → 5 → 16 → 8 → 4 → 2 → 1
//
//It can be seen that this sequence (starting at 13 and finishing at 1) contains 10 terms. Although it has not been proved yet (Collatz Problem), it is thought that all starting numbers finish at 1.
//
//Which starting number, under one million, produces the longest chain?
//
//NOTE: Once the chain starts the terms are allowed to go above one million.

package problems

func LongestCollatzSequence(limit int) int {
	longestLength := 0
	startingNumber := 0
	cache := make(map[int]int)

	for i := 1; i < limit; i++ {
		n := i
		length := 0

		for n != 1 {
			if val, exists := cache[n]; exists {
				length += val
				break
			}

			if n%2 == 0 {
				n = n / 2
			} else {
				n = 3*n + 1
			}
			length++
		}

		cache[i] = length

		if length > longestLength {
			longestLength = length
			startingNumber = i
		}
	}

	return startingNumber
}