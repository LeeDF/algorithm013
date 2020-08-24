package Week_04

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	left := 0
	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])
	right := m*n - 1
	for left <= right {
		mid := left + (right-left)/2
		v := matrix[mid / n][mid % n]
		if v == target{
			return true
		}
		if target < v{
			right = mid-1
		}
		if target > v{
			left = mid+1
		}
	}
	return false
}

func main() {
	matrix := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 50},
	}
	fmt.Println(searchMatrix(matrix, 23))
}
