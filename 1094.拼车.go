func carPooling(trips [][]int, capacity int) bool {
	result := make([]int, 1001)

	result[0] = capacity
	for i := 0; i < 1001; i ++ {
		notFound := true
		for _, x := range trips {
			if x[2] == i {
				if notFound {
					result[i] = result[i-1] + x[0]
				} else {
					result[i] = result[i] + x[0]
				}

				notFound = false
			}

			if x[1] == i {
				if i == 0 {
					result[i] = capacity - x[0]
				} else {
					if notFound {
						result[i] = result[i-1] - x[0]
					} else {
						result[i] = result[i] - x[0]
					}
				}

				notFound = false
			}

			if notFound {
				if i != 0 {
					result[i] = result[i-1]
				}
			}
		}

		if result[i] < 0 {
			return false
		}
	}

	return true
}
