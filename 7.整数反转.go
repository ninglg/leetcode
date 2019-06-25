func reverse(x int) int {
	var y, s int
	var t []int
	var m int
	if x < 0 {
		s = -1
		x = -x
	} else {
		s = 1
	}
	for x != 0 {
		m = x % 10
		t = append(t, m)
		x = x / 10
	}

	for i := 0; i < len(t); i++ {
		y = y*10 + t[i]
	}
    
    if y * s < -2147483648 || y * s > 2147483647 {
		return 0
	}
    
	return y * s
}
