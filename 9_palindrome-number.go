func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	y := x
	var t []int
	for y != 0 {
		t = append(t, y%10)
		y = y / 10
	}

	var r int
	for i := 0; i < len(t); i++ {
		r = r * 10 + t[i]
	}

	if r != x {
		return false
	}

	return true
}
