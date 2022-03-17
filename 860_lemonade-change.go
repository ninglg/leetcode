func lemonadeChange(bills []int) bool {
    var fives, tens int
    for _, x := range bills {
        if x == 5 {
            fives++
        } else if x == 10 {
            tens++
            if fives > 0 {
                fives--
            } else {
                return false
            }
        } else if x == 20 {
            if tens > 0 && fives > 0 {
                tens--
                fives--
            } else if fives >= 3 {
                fives -= 3
            } else {
                return false
            }
        }
    }
    
    return true
}
