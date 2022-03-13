func mostWordsFound(sentences []string) int {
    var ret,max int

    for _, v := range sentences {
        max = len(strings.Split(v, " "))
        if max > ret {
            ret = max
        }
    }

    return ret
}
