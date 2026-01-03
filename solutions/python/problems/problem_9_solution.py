# Problem 9: Special Pythagorean triplet
# A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,
# a^2 + b^2 = c^2.
# For example, 3^2 + 4^2 = 9 + 16 = 25 = 5^2.
# There exists exactly one Pythagorean triplet for which a + b + c = 1000.
# Find the product abc. 

#Link: https://projecteuler.net/problem=9

# Solution:
def special_pythagorean_triplet(n):
    for a in range(1, n // 3): # a must be less than n/3
        for b in range(a + 1, n // 2): # b must be less than n/2
            c = n - a - b # since a + b + c = n
            if a * a + b * b == c * c: # check for Pythagorean triplet
                return a * b * c # return the product abc
    return None # if no triplet is found