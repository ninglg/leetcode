func plusOne(digits []int) []int {
	var result []int

	plus := false
	for i := len(digits)-1; i >= 0; i-- {
		if plus {
			if digits[i] == 9 {
				result = append(result, 0)
				plus = true
			} else {
				result = append(result, digits[i]+1)
				plus = false
			}
		} else {
			if i == len(digits)-1 {
				if digits[len(digits)-1] == 9 {
					result = append(result, 0)
					plus = true
				} else {
					result = append(result, digits[i]+1)
					plus = false
				}
			} else {
				result = append(result, digits[i])
				plus = false
			}
		}
	}
    
    if plus {
		result = append(result, 1)
	}
    
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return result
}
