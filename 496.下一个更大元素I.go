func nextGreaterElement(nums1 []int, nums2 []int) []int {
	var result []int

	for i := 0; i < len(nums1); i++ {
		temp := -1
		flag := false
		for j := 0; j < len(nums2); j++ {
			if flag && nums2[j] > nums1[i] {
				temp = nums2[j]
				break
			}

			if nums2[j] == nums1[i] {
				flag = true
			}
		}

		result = append(result, temp)
	}

	return result
}

