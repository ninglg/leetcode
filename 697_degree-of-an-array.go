func findShortestSubArray(nums []int) int {
    m := make(map[int]int)

    // 先统计各个数字的出现次数，并记录最多的次数
    max := 0
    for _, v := range nums {
        m[v]++
        if m[v] > max {
            max = m[v]
        }
    }

    // 找出出现次数达到最多的可能多个数字
    var n []int
    for k, v := range m {
        if v == max {
            n = append(n, k)
        }
    }

    // 分别找出多个数字的最短子数组长度，以最小的长度返回
    min := len(nums)
    r := make(map[int]int)
    for _, v := range n {
        start, end := 0, 0
        var flag bool
        for i := 0; i < len(nums); i++ {
            if nums[i] == v {
                if !flag {
                    start = i
                    flag = true
                }
                end = i
            }
        }

        r[v] = end - start + 1

        if r[v] < min {
            min = r[v]
        }
    }

    return min
}

