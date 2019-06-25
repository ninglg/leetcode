func isPerfectSquare(num int) bool {
    for i := 0; i <= (num/2 + 1); i++ {
        if i * i == num {
            return true
        } else if i*i > num {
            return false
        }
    }
    
    return false
}
