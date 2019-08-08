func longestCommonPrefix(strs []string) string {
	n := len(strs)
	if n == 0 {
		return ""
	}

	min_len := len(strs[0])
	tmp := 0

	for i := 1; i < n; i++ {
		tmp = len(strs[i])
		if tmp < min_len {
			min_len = len(strs[i])
		}
	}

	ans := ""
	for j := 0; j < min_len; j++ {
		for k := 0; k < n-1; k++ {
			if strs[k][j] != strs[k+1][j] {
				return ans
			}
		}
		ans += string(strs[0][j])
	}
	return ans
}

