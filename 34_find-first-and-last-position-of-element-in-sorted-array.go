func searchRange(nums []int, target int) []int {
    ret := []int{-1, -1}
    var flag bool

    for i := 0; i < len(nums); i++ {
        if nums[i] > target {
            break
        }

        if nums[i] == target {
            if !flag {
                ret[0] = i
                flag = true
            }
            ret[1] = i
        }
    }

    return ret
}

