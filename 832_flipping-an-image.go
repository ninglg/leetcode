func flipAndInvertImage(A [][]int) [][]int {
	var result [][]int
	var t []int
	for i := 0; i < len(A); i++ {
		t = []int{}
		for j := 0; j < len(A[i]); j++ {
			if A[i][len(A[i])-j-1] == 0 {
				t = append(t, 1)
			} else {
				t = append(t, 0)
			}
		}
		result = append(result, t)
	}
	return result
}
