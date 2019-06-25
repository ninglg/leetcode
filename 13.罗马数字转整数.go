func romanToInt(s string) int {
	j := 0
	fi, fx, fc := false, false, false

	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'I':
			if (i+1 < len(s)) && (s[i+1] == 'V' || s[i+1] == 'X') {
				fi = true
			} else {
				j += 1
			}
		case 'V':
			if fi {
				fi = false
				j += 4
			} else {
				j += 5
			}
		case 'X':
			if fi {
				fi = false
				j += 9
			} else {
				if (i+1 < len(s)) && (s[i+1] == 'L' || s[i+1] == 'C') {
					fx = true
				} else {
					j += 10
				}
			}
		case 'L':
			if fx {
				fx = false
				j += 40
			} else {
				j += 50
			}
		case 'C':
			if fx {
				fx = false
				j += 90
			} else {
				if (i+1 < len(s)) && (s[i+1] == 'D' || s[i+1] == 'M') {
					fc = true
				} else {
					j += 100
				}
			}
		case 'D':
			if fc {
				fc = false
				j += 400
			} else {
				j += 500
			}
		case 'M':
			if fc {
				fc = false
				j += 900
			} else {
				j += 1000
			}
		}
	}

	return j
}
