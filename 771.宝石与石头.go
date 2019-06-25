func numJewelsInStones(J string, S string) int {
	var count int
	JJ := []rune(J)
	SS := []rune(S)

	for i := 0; i < len(SS); i++ {
		for j := 0; j < len(JJ); j++ {
			if SS[i] == JJ[j] {
				count++
				break
			}
		}
	}

	return count
}
