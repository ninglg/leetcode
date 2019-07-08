func maximumProduct(nums []int) int {
    sort.Ints(nums)
	length := len(nums)
	a := nums[length-1] * nums[length-2] * nums[length-3]
	b := nums[0] * nums[1] * nums[length-1]
	if a > b {
		return a
	} else {
		return b
	}
}
