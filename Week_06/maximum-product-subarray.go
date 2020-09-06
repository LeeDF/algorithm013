package Week_06

import "fmt"



func maxProduct(nums []int) int {
	ans := nums[0]
	premax := nums[0]
	premin := nums[0]

	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}

	for i:=1;i< len(nums);i++{
		mx ,mn := premax,premin
		premax = max(mx*nums[i], max(mn*nums[i], nums[i]))
		premin = min(mx*nums[i], min(mn*nums[i], nums[i]))
		ans = max(ans, premax)
	}
	return ans
}



func main() {
	fmt.Println(maxProduct([]int{-4, -3, -2}))
}
