func findMaxConsecutiveOnes(nums []int) int {
    max := 0
    j := 0
    
    for _, v := range nums {
        if v == 1 {
            j++
        } else if v == 0 {
            if j > max {
                max = j
            }

            j = 0
        }
    }

    if j > max {
        max = j
    }

    return max
}
