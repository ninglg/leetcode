func repeatedNTimes(A []int) int {
    result :=make(map[int]int)
    for _, x := range A {
        result[x]++
    }
    
    for x, y := range result {
        if y == len(A)/2 {
            return x
        }
    }
    
    return 0
}
