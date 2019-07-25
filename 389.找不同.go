func findTheDifference(s string, t string) byte {
    ss := []byte(s)
    tt := []byte(t)
    var b byte
    
    ttt := make(map[byte]int)
    for _, x := range tt {
        ttt[x]++
    }
    
    for _, x := range ss {
        ttt[x]--
    }
    
    for b, c := range ttt {
        if c == 1 {
            return b
        }
    }
    
    return b
}
