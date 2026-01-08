const problem_1 = require("./problems/problem_1_solution"); //Multiples of 3 and 5
const problem_2 = require("./problems/problem_2_solution"); //Even Fibonacci numbers
const problem_3 = require("./problems/problem_3_solution"); //Largest prime factor
const problem_4 = require("./problems/problem_4_solution"); //Largest palindrome product
const problem_5 = require("./problems/problem_5_solution"); //Smallest multiple
const problem_6 = require("./problems/problem_6_solution"); //Sum square difference
const problem_7 = require("./problems/problem_7_solution"); //10001st prime
const problem_8 = require("./problems/problem_8_solution"); //Largest product in a series
const problem_9 = require("./problems/problem_9_solution"); //Special Pythagorean triplet
const problem_10 = require("./problems/problem_10_solution"); //Sum of primes below two million
const problem_11 = require("./problems/problem_11_solution"); //Largest product in a grid
const problem_12 = require("./problems/problem_12_solution"); //Highly divisible triangular number
const problem_13 = require("./problems/problem_13_solution"); //Large sum


const PROBLEMS = {
  1: problem_1, // Problem 1: Multiples of 3 and 5
  2: problem_2, // Problem 2: Even Fibonacci numbers
  3: problem_3, // Problem 3: Largest prime factor
  4: problem_4, // Problem 4: Largest palindrome product
  5: problem_5, // Problem 5: Smallest multiple
  6: problem_6, // Problem 6: Sum square difference
  7: problem_7, // Problem 7: 10001st prime
  8: problem_8, // Problem 8: Largest product in a series
  9: problem_9, // Problem 9: Special Pythagorean triplet
  10: problem_10, // Problem 10: Sum of primes below two million
  11: problem_11, // Problem 11: Largest product in a grid
  12: problem_12, // Problem 12: Highly divisible triangular number
  13: problem_13, // Problem 13: Large sum
};

module.exports = { PROBLEMS};