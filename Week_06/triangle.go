package Week_06

import "fmt"

// 最后一行就是自己的值
// dp[i][j] = 2 + min (dp[i+1][j], dp[i+1][j+1])

func minimumTotal(triangle [][]int) int {
	m := len(triangle)
	for i := m - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			triangle[i][j] += min(triangle[i+1][j], triangle[i+1][j+1])
		}
	}
	return triangle[0][0]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	fmt.Println(minimumTotal([][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}))
}
