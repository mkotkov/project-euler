// Check if a number is a palindrome
const isPalindrome = (num) => {
    const str = num.toString();
    const reversed = str.split('').reverse().join('');
    return str === reversed;
};

module.exports = isPalindrome;