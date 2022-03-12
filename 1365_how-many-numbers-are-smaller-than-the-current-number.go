func smallerNumbersThanCurrent(nums []int) []int {
	newNums := make([]int, len(nums))
	copy(newNums, nums)
	sort.Ints(newNums)

	m := make(map[int]int)
	for i, v := range newNums {
		if _, ok := m[v]; !ok {
			m[v] = i
		}
	}

	var ret []int
	for _, v := range nums {
		ret = append(ret, m[v])
	}

	return ret
}

