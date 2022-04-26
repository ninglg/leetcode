func canPlaceFlowers(flowerbed []int, n int) bool {
    if n == 0 {
        return true
    }
    
    for i := 0; i < len(flowerbed); i++ {
        if flowerbed[i] == 0 {
            if (i < 1 || flowerbed[i-1] ==0) && (i > (len(flowerbed)-2) || flowerbed[i+1] == 0 ) {
                flowerbed[i] = 1
                n--
                if n == 0 {
                    return true
                }
            }
        }
    }

    return false
}
