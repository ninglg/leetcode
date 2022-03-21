func judgeCircle(moves string) bool {
    // U/D, L/R
    m := [2]int{0, 0}

    for i := 0; i < len(moves); i++ {
        switch moves[i] {
            case 'U':
            m[0]++
            case 'D':
            m[0]--
            case 'L':
            m[1]++
            case 'R':
            m[1]--
        }
    }

    if m[0] == 0 && m[1] == 0 {
        return true
    }

    return false
}
