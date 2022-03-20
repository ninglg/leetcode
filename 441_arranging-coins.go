// 根据数列求和公式推导，先模糊取最大可能值，然后逆序判断精确值。执行速度快。
func arrangeCoins(n int) int {
    i := int(math.Sqrt(float64(2*n)))

    for ; i > 1; i-- {
        if i*(i+1) <= 2*n {
            return i
        }
    }

    return 1
}

