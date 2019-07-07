func checkPerfectNumber(num int) bool {
    s := int(math.Sqrt(float64(num)))
    result := 1
    if num == 0 || num == 1 {
        return false
    }
    
    for i := 2; i <= s; i++ {
		if num % i == 0 {
			result += i
            if i != 1 {
                result += num / i
            }
		}
	}
    
	if result == num {
		return true
	}

	return false
}
