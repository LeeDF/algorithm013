package Week_06

func rob(nums []int) int {
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	var (
		p1  int
		p2  int
		ans int
	)
	for i := range nums {
		p := max(p2, p1+nums[i])
		ans = max(ans, p)
		p1, p2 = p2, p
	}
	return ans
}
