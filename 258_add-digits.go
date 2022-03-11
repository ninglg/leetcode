func addDigits(num int) int {
    for num > 9 {
        num = num % 9
        if num == 0 {
            return 9
        }
    }

    return num
}
