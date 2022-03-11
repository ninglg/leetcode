func sumOfUnique(nums []int) int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}

	var ret int
	for k, v := range m {
		if v == 1 {
			ret += k
		}
	}

	return ret
}
