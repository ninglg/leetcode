func findRelativeRanks(nums []int) []string {
	var result []string
	temp := make([]int, len(nums))
	copy(temp, nums)
	sort.Ints(temp)

	for i := 0; i < len(nums); i++ {
		if nums[i] == temp[len(temp)-1] {
			result = append(result, "Gold Medal")
		} else if nums[i] == temp[len(temp)-2] {
			result = append(result, "Silver Medal")
		} else if nums[i] == temp[len(temp)-3] {
			result = append(result, "Bronze Medal")
		} else {
			for j := 0; j < len(temp)-3; j++ {
				if nums[i] == temp[j] {
					result = append(result, strconv.Itoa(len(temp)-j))
				}
			}
		}
	}
	
	return result
}

