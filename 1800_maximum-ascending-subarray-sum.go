func maxAscendingSum(nums []int) int {
	max, sum := nums[0], nums[0]

	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] > nums[i] {
			sum += nums[i+1]
		} else {
			if sum > max {
				max = sum
			}

			sum = nums[i+1]
		}
	}

	if sum > max {
		max = sum
	}

	return max
}

