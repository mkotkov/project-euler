
// Problem 4: Largest palindrome product
// A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
// Find the largest palindrome made from the product of two 3-digit numbers.

// Link: https://projecteuler.net/problem=4

// Solution:
const isPalindrome = require("../base/is_palindrome");

const largestPalindrome = (n_digits) => {
    let maxPalindrome = 0;
    const lengthLimit = Math.pow(10, n_digits);

    for (let i = lengthLimit - 1; i >= lengthLimit / 10; i--) {
        for (let j = lengthLimit - 1; j >= lengthLimit / 10; j--) {
            const product = i * j;
            if (isPalindrome(product) && product > maxPalindrome) {
                maxPalindrome = product;
            }
        }
    }
    return maxPalindrome;
};

module.exports = largestPalindrome;