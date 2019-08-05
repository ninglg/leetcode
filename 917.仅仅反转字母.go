func reverseOnlyLetters(S string) string {
	if len(S) == 0 {
		return S
	}

	t := []rune(S)
	i,j := 0, len(t)-1
	LOOP1:
		if unicode.IsLetter(t[i]) {
			goto LOOP2
		} else {
			i++
			if i < len(t)-1 {
				goto LOOP1
			}
		}

	LOOP2:
		if unicode.IsLetter(t[j]) {
			goto LOOP3
		} else {
			j--
			if j > 0 {
				goto LOOP2
			}
		}

	LOOP3:
		if i < j {
			t[i], t[j] = t[j], t[i]
			i++
			j--
			goto LOOP1
		}

	return string(t)
}
