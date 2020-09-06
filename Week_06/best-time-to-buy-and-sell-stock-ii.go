package Week_06

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	max := func(x, y int) int {
		if x > y {return x}
		return y
	}

	dp := make([][2]int, len(prices))
	dp[0][0] = 0
	dp[0][1] = 0 - prices[0]
	for i := 1; i < len(prices); i++{
		dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0] - prices[i])
	}
	return dp[len(prices)-1][0]
}
//0 不持有， 1持有
//dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i]) 卖出
//dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i]) 买入

func main()  {
	maxProfit([]int{7,1,5,3,6,4})
}