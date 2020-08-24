package Week_04

import "fmt"

//判断左右哪侧是递增的， 通过递增的区间判断是否在递增区间， 否在在另一侧
func search(nums []int, target int) int {
	left := 0
	right := len(nums)-1
	for left <= right{
		mid := left + (right - left) / 2
		if nums[mid] == target{
			return mid
		}
		if nums[left] == target{
			return left
		}
		if nums[right] == target{
			return right
		}
		if nums[left] < nums[mid] {
			//左边是递增的
			if target > nums[left] && target < nums[mid]{
				right = mid-1
			}else{
				left = mid +1
			}
		}else{
			//右边是递增的
			if target > nums[mid] && target < nums[right]{
				left = mid +1
			}else{
				right = mid-1

			}
		}
	}
	return -1
}

func main()  {
	fmt.Println(search([]int{4,5,6,7,0,1,2}, 1))
}