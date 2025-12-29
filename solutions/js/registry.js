const problem_1 = require("./problems/problem_1_solution"); //Multiples of 3 and 5
const problem_2 = require("./problems/problem_2_solution"); //Even Fibonacci numbers
const problem_3 = require("./problems/problem_3_solution"); //Largest prime factor
const problem_4 = require("./problems/problem_4_solution"); //Largest palindrome product
const problem_5 = require("./problems/problem_5_solution"); //Smallest multiple

const PROBLEMS = {
  1: problem_1, // Problem 1: Multiples of 3 and 5
  2: problem_2, // Problem 2: Even Fibonacci numbers
  3: problem_3, // Problem 3: Largest prime factor
  4: problem_4, // Problem 4: Largest palindrome product
  5: problem_5, // Problem 5: Smallest multiple
};

module.exports = { PROBLEMS};