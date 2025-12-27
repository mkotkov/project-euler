// Problem 3: Largest prime factor
// The prime factors of 13195 are 5, 7, 13 and 29.
// What is the largest prime factor of the number 600851475143?  

//  Link: https://projecteuler.net/problem=3

// Solution:
const isPrime = require("../base/is_prime");

const largestPrimeFactor = (n) => {
    let factor = 2;
    while (factor * factor <= n) {
        if (n % factor === 0 && isPrime(factor)) {
            n /= factor;
        } else {
            factor++;
        }
    }
    return n;
};

module.exports = largestPrimeFactor;