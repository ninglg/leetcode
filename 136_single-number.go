func singleNumber(nums []int) int {
    var i int
	for _, x := range nums {
		i = i^x
	}
	return i
}
