func twoSum(numbers []int, target int) []int {
    var result []int
    
    for i := 0; i < len(numbers)-1; i++ {
        for j := i+1; j < len(numbers); j++ {
            temp := numbers[i] + numbers[j]
            if temp == target {
                result = append(result, []int{i+1, j+1}...)
                return result
            } else if temp > target {
                break
            }
        }
    }
    
    return nil
}
