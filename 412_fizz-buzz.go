func fizzBuzz(n int) []string {
    var result []string
    var f3, f5 bool
    
    for i := 1; i <= n; i++ {
        f3, f5 = false, false
        
        if i % 3 == 0 {
            f3 = true
        }
        
        if i % 5 == 0 {
            f5 = true
        }
        
        if f3 && f5 {
            result = append(result, "FizzBuzz")
        } else if f3 {
            result = append(result, "Fizz")
        } else if f5 {
            result = append(result, "Buzz")
        } else {
            result = append(result, strconv.Itoa(i))
        }
    }
    
    return result
}

