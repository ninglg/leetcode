func missingNumber(nums []int) int {
	result := make([]int, len(nums)+1)

	for _, x := range nums {
		result[x] = 1
	}

	for y, z := range result {
		if z == 0 {
			return y
		}
	}
	
	return 0
}
