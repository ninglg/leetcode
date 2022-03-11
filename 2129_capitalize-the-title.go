func capitalizeTitle(title string) string {
	sc := strings.Split(title, " ")

	for i := 0; i < len(sc); i++ {
		if len(sc[i]) <= 2 {
			sc[i] = strings.ToLower(sc[i])
		} else {
			sc[i] = strings.ToUpper(sc[i][:1]) + strings.ToLower(sc[i][1:])
		}
	}

	return strings.Join(sc, " ")
}

