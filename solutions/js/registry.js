const problem_1 = require("./problems/problem_1_solution"); //Multiples of 3 and 5
const problem_2 = require("./problems/problem_2_solution"); //Even Fibonacci numbers
const problem_3 = require("./problems/problem_3_solution"); //Largest prime factor

const PROBLEMS = {
  1: problem_1, // Problem 1: Multiples of 3 and 5
  2: problem_2, // Problem 2: Even Fibonacci numbers
  3: problem_3, // Problem 3: Largest prime factor  
};

module.exports = { PROBLEMS};