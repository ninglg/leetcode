func findLengthOfLCIS(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    
    y, m := 1, 1
    for i := 1; i < len(nums); i++ {
        if nums[i] > nums[i-1] {
            y++
        } else {
            y = 1
        }
        
        if y > m {
            m = y
        }
    }
        
    return m
}
