# Problem 4: Largest palindrome product

# A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
# Find the largest palindrome made from the product of two 3-digit numbers.

# Solution
from base.is_palindrome import is_palindrome

def largest_palindrome_product(digits):
    max_palindrome = 0
    lower_limit = 10**(digits - 1)
    upper_limit = 10**digits

    for i in range(upper_limit - 1, lower_limit - 1, -1):
        for j in range(i, lower_limit - 1, -1):
            product = i * j
            if product <= max_palindrome:
                break  # Products will decrease after this point
            if is_palindrome(product):
                max_palindrome = product

    return max_palindrome