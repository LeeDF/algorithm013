package Week_06

import "fmt"

func longestCommonSubsequence(text1 string, text2 string) int {
	n1 := len(text1)
	n2 := len(text2)
	dp := make([][]int, n1+1)
	for i := 0;i <= n1; i++{
		dp[i] = make([]int, n2+1)
	}
	for i := 1;i <= n1; i++{
		for j := 1;j <= n2; j++{
			if text1[i-1] == text2[j-1]{
				dp[i][j] = 1 + dp[i-1][j-1]
			}else{
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[n1][n2]
}

func max(x,y int) int {
	if x >= y{
		return x
	}
	return y
}

func main()  {
	fmt.Println(longestCommonSubsequence("abcde", "ace"))
}
