func pivotIndex(nums []int) int {
    all := 0

    for _, v := range nums {
        all += v
    }

    sum := 0
    for i := 0; i < len(nums); i++ {
        if all == (sum * 2 + nums[i]) {
            return i
        }

        sum += nums[i]
    }

    return -1
}
