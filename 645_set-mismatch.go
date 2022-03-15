func findErrorNums(nums []int) []int {
    m := make(map[int]int, len(nums))

    for _, v := range nums {
        m[v]++
    }

    var repeat,miss int
    for i := 1; i < len(nums)+1; i++ {
        if m[i] == 2 {
            repeat = i
        }

        if m[i] == 0 {
            miss = i
        }

        if repeat != 0 && miss != 0 {
            break
        }
    }

    ret := []int{repeat, miss}
    return ret
}

