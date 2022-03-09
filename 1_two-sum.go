func twoSum(nums []int, target int) []int {
	h := make(map[int]int)

	for i, value := range nums {
		if wanted, ok := h[value]; ok {
			return []int{wanted, i}
		} else {
			h[target-value] = i
		}
	}

	return nil
}
