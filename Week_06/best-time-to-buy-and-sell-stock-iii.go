package Week_06

import (
	"fmt"
	"strconv"
)

func maxProfitiii(prices []int) int {
	if len(prices) == 0{
		return 0
	}
	max := func(x, y int) int {
		if x > y {return x}
		return y
	}
	strconv.Itoa()
	dp := make([][3][2]int, len(prices))
	dp[0][0] = [2]int{0,-prices[0]}
	for i:=1 ;i < len(prices); i++ {
		if i > 1{
			dp[i][2][0] = max(dp[i-1][2][0], dp[i-1][2][1]+prices[i])
			dp[i][2][1] = max(dp[i-1][2][1], dp[i-1][1][0]-prices[i])
		}

		dp[i][1][0] = max(dp[i-1][1][0], dp[i-1][1][1]+prices[i])
		dp[i][1][1] = max(dp[i-1][1][1], -prices[i])

	}
	return dp[len(prices)-1][2][0]
}

//0 不持有， 1持有
//k 交易次数
//dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1]+prices[i]) 卖出
//dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0]-prices[i]) 买入

func main()  {
	fmt.Println(maxProfitiii([]int{1,2,3,4,5}))
}