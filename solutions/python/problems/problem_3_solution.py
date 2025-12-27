# Problem 3: Largest prime factor
# The prime factors of 13195 are 5, 7, 13 and 29.
# What is the largest prime factor of the number 600851475143?  

#  Link: https://projecteuler.net/problem=3

# Solution:
from base.is_prime import num_is_prime as is_prime

def largest_prime_factor(n):
   
    largest_factor = None
    # Check for number of 2s that divide n
    while n % 2 == 0:
        largest_factor = 2
        n //= 2

    # n must be odd at this point, so we can skip even numbers
    for i in range(3, int(n**0.5) + 1, 2):
        while n % i == 0 and is_prime(i):
            largest_factor = i
            n //= i

    # This condition is to check if n is a prime number greater than 2
    if n > 2 and is_prime(n):
        largest_factor = n

    return largest_factor