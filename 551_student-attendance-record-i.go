func checkRecord(s string) bool {
    var aCount, lCount int
    var lFlag bool
    for i := 0; i < len(s); i++ {
        if s[i] == 'A' {
            aCount++
            if aCount == 2 {
                return false
            }
        }

        if s[i] == 'L' {
            if !lFlag {
                lCount = 1
            } else {
                lCount++
            }

            if lCount == 3 {
                return false
            }

            lFlag = true
        } else {
            lFlag = false
        }
    }

    return true
}
