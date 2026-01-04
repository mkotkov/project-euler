# Problem 10: Sum of primes below two million
# The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.
# Find the sum of all the primes below two million.

# Link: https://projecteuler.net/problem=10

# Solution:
from base.is_prime import num_is_prime as is_prime

def sum_of_primes_below(n): # sum of all primes below n
    total = 0 # initialize total sum
    for num in range(2, n): # iterate through all numbers below n
        if is_prime(num): # check if the number is prime
            total += num # add prime number to total
    return total # return the total sum of primes below n