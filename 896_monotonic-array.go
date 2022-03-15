func isMonotonic(A []int) bool {
	var big, small bool
	for i := 1; i < len(A); i++ {
		if !small && !big {
			if A[i] > A[i-1] {
				big = true
			}
			
			if A[i] < A[i-1] {
				small =true
			}
	}
		if big && (A[i] < A[i-1]) {
			return false
		}

		if small && (A[i] > A[i-1]) {
			return false
		}

	}

	return true
}
