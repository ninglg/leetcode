func relativeSortArray(arr1 []int, arr2 []int) []int {
	var arr3 []int
	arr4 := make([]int, len(arr1))
	copy(arr4, arr1)

	for _, x := range arr2 {
		for i := 0; i < len(arr1); i++ {
			if arr1[i] == x {
				arr3 = append(arr3, x)
			}
		}

		for j := 0; j < len(arr4); j++ {
			if arr4[j] == x {
				arr4 = append(arr4[:j], arr4[j+1:]...)
				j = -1
			}
		}
	}

	sort.Ints(arr4)
	arr3 = append(arr3, arr4...)

	return arr3
}

