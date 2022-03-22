func containsDuplicate(nums []int) bool {
    m := make(map[int]int)

    for _, v := range nums {
        m[v]++
        if m[v] >= 2 {
            return true
        }
    }

    return false
}
