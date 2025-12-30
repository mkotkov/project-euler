from problems.problem_1_solution import multiples_of_3_and_5    #Multiples of 3 and 5
from problems.problem_2_solution import fibonacci_even_sum      #Even Fibonacci numbers
from problems.problem_3_solution import largest_prime_factor    #Largest prime factor
from problems.problem_4_solution import largest_palindrome_product    #Largest palindrome product
from problems.problem_5_solution import SmallestMultiple as smallest_multiple         #Smallest multiple
from problems.problem_6_solution import sum_square_difference #Sum square difference

problem_registry = {
    1: multiples_of_3_and_5,    #Multiples of 3 and 5
    2: fibonacci_even_sum,      #Even Fibonacci numbers
    3: largest_prime_factor,     #Largest prime factor
    4: largest_palindrome_product, #Largest palindrome product
    5: smallest_multiple,        #Smallest multiple
    6: sum_square_difference     #Sum square difference
}