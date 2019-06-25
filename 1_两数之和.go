func twoSum(nums []int, target int) []int {
    result := []int{}
    count := len(nums)
    
    for i := 0; i < count - 1; i++ {
        for j := i+1; j < count; j++ {
            if nums[i] + nums[j] == target {
                result = append(result, i)
                result = append(result, j)
                return result
            }
        }
    }
    return nil
}
