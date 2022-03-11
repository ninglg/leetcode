func hammingWeight(num uint32) int {
    var ret int

    for num != 0 {
        if num & 1 == 1 {
            ret++
        }

        num = num >> 1
    }

    return ret
}
