def fibonacci_sum(limit):
    a, b = 0, 1
    sum_even = 0
    while b < limit:
        if b%2 == 0:
            sum_even += b
        a, b = b, a + b
    return sum_even