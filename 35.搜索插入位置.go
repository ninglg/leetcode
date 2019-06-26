func searchInsert(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			return i
		} else if nums[i] > target {
			if i > 0 {
				return i
			}
			return 0
		}
	}
	return len(nums)
}
