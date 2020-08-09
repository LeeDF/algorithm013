package Week_02

//https://leetcode-cn.com/problems/chou-shu-lcof/
//丑数等于丑数乘以丑数
//2,3,5乘以已经生成好且排好序的丑数,取最小的作为下一个丑数,用三个指针记录分别乘以到哪个元素了
func nthUglyNumber(n int) int {
	ret := make([]int, 0, n)
	ret = append(ret, 1)
	var p2, p3, p5 int
	for i := 1; i < n; i++ {
		ret = append(ret, min(min(ret[p2]*2, ret[p3]*3), ret[p5]*5))
		if ret[p2]*2 == ret[i] {
			//表示最小的是p2
			p2++
		}
		if ret[p3]*3 == ret[i] {
			p3++
		}
		if ret[p5]*5 == ret[i] {
			p5++
		}
	}
	return ret[n-1]

}

func min(x int, y int) int {
	if x < y {
		return x
	}
	return y
}
