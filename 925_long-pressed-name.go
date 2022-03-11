func isLongPressedName(name string, typed string) bool {
	n := []rune(name)
	t := []rune(typed)

	nn := []map[rune]int{}
	tt := []map[rune]int{}

	for i := 0; i < len(n); i++ {
		if i > 0 && len(nn) > 0 {
			if _, ok := nn[len(nn)-1][n[i]]; ok {
				nn[len(nn)-1][n[i]]++
			} else {
				nn = append(nn, map[rune]int{n[i]:1})
			}
		} else {
			nn = append(nn, map[rune]int{n[i]:1})
		}
	}

	for i := 0; i < len(t); i++ {
		if i > 0 && len(tt) > 0 {
			if _, ok := tt[len(tt)-1][t[i]]; ok {
				tt[len(tt)-1][t[i]]++
			} else {
				tt = append(tt, map[rune]int{t[i]:1})
			}
		} else {
			tt = append(tt, map[rune]int{t[i]:1})
		}
	}

	if len(nn) != len(tt) {
		return false
	}

	for i := 0; i < len(nn); i++ {
		nnn := nn[i]
		ttt := tt[i]
		for a, b := range nnn {
			for c, d := range ttt {
				if (a != c) || (b > d) {
					return false
				}
			}
		}
	}

	return true
}
