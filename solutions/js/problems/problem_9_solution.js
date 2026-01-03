// Problem 9: Special Pythagorean triplet
// A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,
// a^2 + b^2 = c^2
//
// For example, 3^2 + 4^2 = 9 + 16 = 25 = 5^2.
//
// There exists exactly one Pythagorean triplet for which a + b + c = 1000.
// Find the product abc.

// https://projecteuler.net/problem=9

// Solution:

function SpecialPythagoreanTriplet(n) {
  for (let a = 1; a < n / 3; a++) {
    for (let b = a + 1; b < n / 2; b++) {
      const c = n - a - b;
      if (a * a + b * b === c * c) {
        return a * b * c;
      }
    }
  }
  return null; // No triplet found
}

module.exports = SpecialPythagoreanTriplet;
