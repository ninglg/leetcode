func largestAltitude(gain []int) int {
    var sum,ret int
    for _, v := range gain {
        sum += v
        if sum > ret {
            ret = sum
        }
    }

    return ret
}
