package Week_04

import "fmt"

func jump(nums []int) int {
	ans := 0
	start := 0
	end := 0
	max := 0
	for end < len(nums)-1 {
		for i := start; i <= end; i++ {
			if nums[i] + i > max{
				max = nums[i] + i
			}
			start = i+1
		}
		end = max
		ans++
	}

	return ans
}

func main() {
	fmt.Println(jump([]int{2,3,1,2,4}))
}
