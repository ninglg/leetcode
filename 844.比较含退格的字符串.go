func backspaceCompare(S string, T string) bool {
    s := []rune(S)
    t := []rune(T)
    
    var ss, tt []rune
    for i := 0; i < len(s); i++ {
        if s[i] == '#' {
            if len(ss) == 1 {
                ss = []rune{}
            }
            
            if len(ss) > 1 {
                ss = ss[:len(ss)-1]
            }
            
            continue
        }

        ss = append(ss, s[i])
    }
    
    for i := 0; i < len(t); i++ {
        if t[i] == '#' {
            if len(tt) == 1 {
                tt = []rune{}
            }
            
            if len(tt) > 1 {
                tt = tt[:len(tt)-1]
            }
            
            continue
        }

        tt = append(tt, t[i])
    }
    
    return string(ss) == string(tt)
}
