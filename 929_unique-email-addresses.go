func numUniqueEmails(emails []string) int {
	result := make(map[string]struct{})
	var s, k []string
	var t, j string
	for _, x := range emails {
		s = strings.Split(x, "@")
		t = s[1]
		k = strings.Split(s[0], "+")
		j = strings.Replace(k[0], ".", "", -1)
		result[j+"@"+t] = struct{}{}
	}

	return len(result)
}
