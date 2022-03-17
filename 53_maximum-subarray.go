func maxSubArray(nums []int) int {
	var k, t int
	t = nums[0]
	for i := 0; i < len(nums); i++ {
		k = k + nums[i]
		if k > t {
			t = k
		}
		if k < 0 {
			k = 0
		}
	}

	return t
}
