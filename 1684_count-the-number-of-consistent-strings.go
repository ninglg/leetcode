func countConsistentStrings(allowed string, words []string) int {
    m := make(map[int32]bool, len(allowed))
    for _, v := range allowed {
        m[v] = true
    }

    var ret int
    var flag bool
    for _, w := range words {
        flag = true
        for _, v := range w {
            if !m[v] {
                flag = false
                break
            }
        }

        if flag {
            ret++
        }
    }

    return ret
}
