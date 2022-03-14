func isValid(s string) bool {
	t := []rune(s)
	var k []rune

	for i := 0; i < len(t); i++ {
		switch t[i] {
		case '(':
			k = append(k, t[i])
		case '[':
			k = append(k, t[i])
		case '{':
			k = append(k, t[i])
		case ')':
			if len(k) == 0 || k[len(k)-1] != '(' {
				return false
			} else {
				k = k[:len(k)-1]
			}
		case ']':
			if len(k) == 0 || k[len(k)-1] != '[' {
				return false
			} else {
				k = k[:len(k)-1]
			}
		case '}':
			if len(k) == 0 || k[len(k)-1] != '{' {
				return false
			} else {
				k = k[:len(k)-1]
			}
		}
	}

	if len(k) == 0 {
		return true
	}
	return false
}
