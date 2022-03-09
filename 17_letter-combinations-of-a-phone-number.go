func letterCombinations(digits string) []string {
	m := make(map[int]string)
	m[2] = "abc"
	m[3] = "def"
	m[4] = "ghi"
	m[5] = "jkl"
	m[6] = "mno"
	m[7] = "pqrs"
	m[8] = "tuv"
	m[9] = "wxyz"

	var ret []string
	for i := 0; i < len(digits); i++ {
		d, _ := strconv.Atoi(string(digits[i]))
		if len(ret) == 0 {
			for k := 0; k < len(m[d]); k++ {
				ret = append(ret, string(m[d][k]))
			}
		} else {
			l := len(ret)
			for j := 0; j < l; j++ {
				for k := 0; k < len(m[d]); k++ {
					ret = append(ret, ret[j]+string(m[d][k]))
				}
			}
			ret = ret[l:]
		}
	}

	return ret
}
