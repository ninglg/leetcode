func lengthOfLastWord(s string) int {
    ss := strings.Fields(s)
    if len(ss) == 0 {
        return 0
    }
    
    return len(ss[len(ss)-1])
}
