func areNumbersAscending(s string) bool {
    sc := strings.Split(s, " ")
    var i int64
    i = -1

    for _, v := range sc {
        vv, err := strconv.ParseInt(v, 10, 64)
        if err != nil {
            continue
        }

        if vv <= i {
            return false
        }
        
        i = vv
    }

    return true
}
