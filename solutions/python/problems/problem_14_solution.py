# Problem 14: Longest Collatz Sequence
# The following iterative sequence is defined for the set of positive integers:
# n → n/2 (n is even)
# n → 3n + 1 (n is odd)
# Using the rule above and starting with 13, we generate the following sequence:
# 13 → 40 → 20 → 10 → 5 → 16 → 8 → 4 → 2 → 1
# It can be seen that this sequence (starting at 13 and finishing at 1) contains 10 terms.
# Although it has not been proved yet (Collatz Problem), it is thought that all starting numbers finish at 1.
# Which starting number, under one million, produces the longest chain? 
# Link to problem: https://projecteuler.net/problem=14

# Solution
def longest_collatz_sequence(limit):
    def collatz_sequence_length(n, memo={}):
        original_n = n
        length = 0
        while n != 1:
            if n in memo:
                length += memo[n]
                break
            if n % 2 == 0:
                n //= 2
            else:
                n = 3 * n + 1
            length += 1
        memo[original_n] = length + 1  # +1 for the last term (1)
        return memo[original_n]

    max_length = 0
    starting_number = 0
    for i in range(1, limit):
        length = collatz_sequence_length(i)
        if length > max_length:
            max_length = length
            starting_number = i

    return starting_number