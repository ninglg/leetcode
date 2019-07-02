func dominantIndex(nums []int) int {
	maxNum := nums[0]
	j := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > maxNum {
			maxNum = nums[i]
			j = i
		}
	}

	for i := 0; i < len(nums); i++ {
		if (nums[i] != maxNum) && (maxNum < 2*nums[i]) {
			return -1
		}
	}

	return j
}
