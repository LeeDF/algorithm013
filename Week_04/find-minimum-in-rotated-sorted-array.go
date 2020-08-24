package Week_04

import "fmt"

func findMin(nums []int) int {
	left := 0
	right := len(nums)-1
	if nums[left] < nums[right]{
		return nums[left]
	}

	for left <= right{
		mid := left + (right-left) / 2
		if mid == left{
			if nums[left] < nums[right]{
				return nums[left]
			}else{
				return nums[right]
			}
		}
		if nums[mid] > nums[mid+1] && nums[mid] > nums[mid-1]{
			return nums[mid+1]
		}
		if nums[mid] < nums[mid+1] && nums[mid] < nums[mid-1]{
			return nums[mid]
		}
		if nums[mid] < nums[mid+1] && nums[mid] > nums[mid-1]{
			if nums[left] < nums[mid]{
				left = mid+1
			}else{
				right = mid-1
			}
		}

	}
	return -1
}


func main()  {
	fmt.Println(findMin([]int{1}))
}
