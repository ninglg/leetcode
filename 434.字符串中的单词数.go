func countSegments(s string) int {
    return len(strings.Fields(strings.TrimSpace(s)))
}
