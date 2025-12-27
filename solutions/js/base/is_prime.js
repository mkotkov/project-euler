// Costome function to check if a number is prime
const IsPrime = (n) => {
	if (n <= 1) {
		return false;
	}
	for (let i = 2; i * i <= n; i++) {
		if (n % i === 0) {
			return false;
		}
	}
	return true;
};

module.exports = IsPrime;