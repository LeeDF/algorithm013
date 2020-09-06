package Week_06

import "fmt"

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	a := make([]int, n)
	if obstacleGrid[0][0] == 0{
		a[0] = 1
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				a[j] = 0
			}else if j >= 1 && obstacleGrid[i][j-1] == 0 {
				a[j] += a[j-1]
			}
		}
	}
	return a[n-1]
}

func main() {
	fmt.Println(uniquePathsWithObstacles([][]int{
		{1, 0},
	}))
}

// a[m][0] = 1, a[0][n] = 1
//
