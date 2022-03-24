// 时间复杂度较高，但占用空间小
func findLHS(nums []int) int {
	max := 0
	for i := 0; i < len(nums)-max; i++ {
		small, big, flag := 1, 1, 1
		for j := i + 1; j < len(nums); j++ {
			if nums[j] == nums[i]-1 {
				small++
				continue
			}

			if nums[j] == nums[i]+1 {
				big++
				continue
			}

			if nums[j] == nums[i] {
				small++
				big++
				flag++
				continue
			}
		}

		if big == flag && small == flag {
			big = 0
			small = 0
		}

		if big < small {
			big = small
		}

		if big > max {
			max = big
		}
	}

	return max
}

