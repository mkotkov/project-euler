// Problem 5: Smallest multiple
// Find the smallest positive number that is evenly divisible by all of the numbers from 1 to 20.

// Solution:
const SmallestMultiple = (end) => {
    let start = 1;
    if (start < 1 || end < start) {
        console.log("Invalid range: numbers must be >= 1 and end >= start");
        return -1;
    };

    let result = 1;
    for (let i = start; i <= end; i++) {
        result = lcm(result, i);
    };
    return result;
};

// Helper function to compute least common multiple (LCM)
const lcm = (a, b) => {
    return a * b / gcd(a, b);
};

// Helper function to compute greatest common divisor (GCD)
const gcd = (a, b) => {
    while (b !== 0) {
        [a, b] = [b, a % b];
    };
    return a;
};

module.exports = SmallestMultiple;