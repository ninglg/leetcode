func topKFrequent(nums []int, k int) []int {
	r := make(map[int]int)
	for _, num := range nums {
		r[num]++
	}

	s := make(map[int][]int)
	for x, y := range r {
		s[y] = append(s[y], x)
	}

	var t []int
	for i := len(nums); i > 0; i-- {
		if _, ok := s[i]; ok {
			t = append(t, s[i]...)
		}

		if k <= len(t) {
			return t
		}
	}

	return t
}
