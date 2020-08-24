package Week_04

import "fmt"

//1.暴力dfs
//2.记录最大能跳多远
func canJump(nums []int) bool {
	k := 0
	for i := range nums {
		if k <= i{
			return false
		}
		k = max(k, i+nums[i])
	}
	return true

}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	fmt.Println(canJump([]int{2, 3, 1, 1, 4}))
}
