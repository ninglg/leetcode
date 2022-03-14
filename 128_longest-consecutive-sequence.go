func longestConsecutive(nums []int) int {
	hashmap := make(map[int]bool)
	for _, i := range nums {
		hashmap[i] = true
	}

	result := 0
	for current := range hashmap {
		if !hashmap[current-1] {
			//不存在比current小的数，则当做序列的首位
			c := current
			cnt := 1
			for hashmap[c+1] {
				c++
				cnt++
			}

			if result < cnt {
				result = cnt
			}
		}
	}

	return result
}

