func majorityElement(nums []int) int {
    result := make(map[int]int)
    t := len(nums)/2
    
    for _, x := range nums {
        result[x]++
        if result[x] > t {
            return x
        }
    }
    
    return 0
}
