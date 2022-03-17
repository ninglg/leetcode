func reverseVowels(s string) string {
    if len(s) == 0 {
		return s
	}
    
	t := []rune(s)
	m := map[rune]struct{}{'a': struct{}{}, 'e': struct{}{}, 'i': struct{}{}, 'o': struct{}{}, 'u': struct{}{}, 'A': struct{}{}, 'E': struct{}{}, 'I': struct{}{}, 'O': struct{}{}, 'U': struct{}{}}

	i,j := 0, len(t)-1
	LOOP1:
		if _, ok := m[t[i]]; ok {
			goto LOOP2
		} else {
			i++
			if i < len(t)-1 {
				goto LOOP1
			}
		}

	LOOP2:
		if _, ok := m[t[j]]; ok {
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

