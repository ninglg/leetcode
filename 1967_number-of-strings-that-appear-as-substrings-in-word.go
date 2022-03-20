func numOfStrings(patterns []string, word string) int {
    count := 0

    for _, v := range patterns {
        if strings.Index(word, v) != -1 {
            count++
        }
    }

    return count
}

