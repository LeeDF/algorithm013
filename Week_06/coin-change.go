package Week_06

import "fmt"

//爬楼梯
//f(i) = min(f(i-1), f(i-2), f(i-5)) + 1
func coinChange(coins []int, amount int) int {

	dp := make([]int, amount+coins[len(coins)-1]+1)

	for i := 0; i <= amount; i++ {
		for _, c := range coins {
			if i > 0 && dp[i] == 0 {
				//说明无法进入此台阶，则跳过
				break
			}
			if dp[c+i] == 0 {
				// 没有值说明没进入过， 则+1， 否在就是之前的值
				dp[c+i] = dp[i] + 1
			}
		}
	}
	if dp[amount] == 0{
		return -1
	}
	return dp[amount]
}

func main()  {
	fmt.Println(coinChange([]int{2147483647,}, 2))
}
