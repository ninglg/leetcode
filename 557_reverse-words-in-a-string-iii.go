func reverseWords(s string) string {
    t := strings.Split(s, " ")

    for i := 0; i < len(t); i++ {
    	y := []rune(t[i])
    	for j, z := 0, len(y)-1; j < z; j,z = j+1,z-1 {
    		y[j], y[z] = y[z], y[j]
    	}
    	t[i] = string(y)
    }
    
    return strings.Join(t, " ")
}
