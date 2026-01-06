from problems.problem_1_solution import multiples_of_3_and_5        #Multiples of 3 and 5
from problems.problem_2_solution import fibonacci_even_sum          #Even Fibonacci numbers
from problems.problem_3_solution import largest_prime_factor        #Largest prime factor
from problems.problem_4_solution import largest_palindrome_product  #Largest palindrome product
from problems.problem_5_solution import smallestMultiple            #Smallest multiple
from problems.problem_6_solution import sum_square_difference       #Sum square difference
from problems.problem_7_solution import nth_prime                   #10001st prime    
from problems.problem_8_solution import largest_product_in_series   #Largest product in a series
from problems.problem_9_solution import special_pythagorean_triplet #Special Pythagorean triplet
from problems.problem_10_solution import sum_of_primes_below        #Sum of primes below two million
from problems.problem_11_solution import largest_product_in_grid  #Largest product in a grid

# Registry mapping problem numbers to their solution functions
problem_registry = {
    1: multiples_of_3_and_5,            #Multiples of 3 and 5
    2: fibonacci_even_sum,              #Even Fibonacci numbers
    3: largest_prime_factor,            #Largest prime factor
    4: largest_palindrome_product,      #Largest palindrome product
    5: smallestMultiple,                #Smallest multiple
    6: sum_square_difference,           #Sum square difference
    7: nth_prime,                       #10001st prime
    8: largest_product_in_series,       #Largest product in a series
    9: special_pythagorean_triplet,      #Special Pythagorean triplet
    10: sum_of_primes_below,             #Sum of primes below two million
    11: largest_product_in_grid         #Largest product in a grid
}