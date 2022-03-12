func maximum69Number(num int) int {
	s := strconv.FormatInt(int64(num), 10)
	s = strings.Replace(s, "6", "9", 1)
	ret, _ := strconv.ParseInt(s, 10, 64)

	return int(ret)
}

