// Problem 10: Sum of primes below two million
// The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.
// Find the sum of all the primes below two million.

// Link: https://projecteuler.net/problem=10

// Solution:
const isPrime = require("../base/is_prime");

// Function to calculate the sum of all primes below a given limit
function sumOfPrimesBelow(limit) {
    let sum = 0; // Initialize sum to 0
    for (let i = 2; i < limit; i++) { // Iterate through all numbers below the limit
        if (isPrime(i)) { // Check if the number is prime
            sum += i; // Add prime number to sum
        }
    }
    return sum;
}

module.exports = sumOfPrimesBelow;