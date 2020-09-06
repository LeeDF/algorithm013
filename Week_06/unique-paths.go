package Week_06

func uniquePaths(m int, n int) int {
	a := make([][]int, n)

	for i := n - 1; i >= 0; i-- {
		ai := make([]int, m)
		a[i] = ai
		for j := m - 1; j >= 0; j-- {
			if i == n-1 || j == m-1{
				a[i][j] = 1
			}else{
				a[i][j] = a[i+1][j] + a[i][j+1]
			}
		}
	}
	return a[0][0]

}

//a. opt[i][j] = opt[i+1][j] + opt[i][j+1]
//b. a[i][j]
//c. f(i,j) = a[i+1][j] + a[i][j+1]

func main() {
	uniquePaths(7, 3)
}
