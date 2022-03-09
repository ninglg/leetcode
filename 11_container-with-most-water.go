func maxArea(height []int) int {
	i := 0
	j := len(height) - 1
	max := 0

	// 双指针解法，因为底部在逐渐向内缩小，只有向内移动height低的一侧才可能出现更大的max
	for i < j {
		ret := 0
		if height[i] < height[j] {
			ret = height[i] * (j - i)
			i++
		} else {
			ret = height[j] * (j - i)
			j--
		}

		if ret > max {
			max = ret
		}
	}

	return max
}