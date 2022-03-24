func isPalindrome(s string) bool {
	b := []rune(strings.ToLower(s))
	var t string
	for _, v := range b {
		if v >= 'a' && v <= 'z' {
			t = t + string(v)
		}

		if v >= '0' && v <= '9' {
			t = t + string(v)
		}
	}

	i, j := 0, len(t)-1
	for i < j {
		if t[i] != t[j] {
			return false
		} else {
			i++
			j--
		}
	}

	return true
}

