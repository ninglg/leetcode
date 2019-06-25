func selfDividingNumbers(left int, right int) []int {
	var result []int
	var t, k int
	var flag bool

	for i := left; i <= right; i++ {
		t = i
		flag = true
		for t != 0 {
			k = t % 10
			t = t / 10
			if k == 0 || i%k != 0 {
				flag = false
				break
			}
		}

		if flag {
			result = append(result, i)
		}
	}

	return result
}
