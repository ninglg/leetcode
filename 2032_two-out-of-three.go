// 通过位运算标记出现的数组，最后判断至少出现在2个数组中的可能值
func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
    m := make(map[int]int)

    for _, v := range nums1 {
        m[v] = 1
    }

    for _, v := range nums2 {
        if _, ok := m[v]; ok {
            m[v] = m[v]|2
        } else {
            m[v] = 2
        }
    }

    for _, v := range nums3 {
        if _, ok := m[v]; ok {
            m[v] = m[v]|4
        } else {
            m[v] = 4
        }
    }

    var ret []int
    for k, v := range m {
        if v == 3 || v == 5 || v == 6 || v== 7 {
            ret = append(ret, k)
        }
    }

    return ret
}
