func removeElement(nums []int, val int) int {
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] != val {
			for j := 0; j < i; j++ {
				if nums[j] == val {
					nums[j] = nums[i]
					nums[i] = val
				}
			}
		}
	}
	for k := 0; k < len(nums); k++ {
		if nums[k] == val {
			return k
		}
	}

	return len(nums)
}
