func isPowerOfFour(n int) bool {
	if n <= 0 {
		return false
	}

	if n == 1 {
		return true
	}

	for n > 0 {
		if (n%4 != 0) && (n != 1) {
			return false
		}

		n = n / 4
	}

	return true
}
