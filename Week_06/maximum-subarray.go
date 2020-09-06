package Week_06

import "fmt"


//dp[i] = max(dp[i], dp[i-1] + dp[i])

func maxSubArray(nums []int) int {
	sum := 0
	ans := nums[0]
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	for i :=  range nums{
		if sum > 0{
			sum += nums[i]
		}else{
			sum = nums[i]
		}
		ans = max(ans, sum)
	}
	return ans
}

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
