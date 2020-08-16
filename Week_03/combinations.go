package Week_03

import "fmt"

//https://leetcode-cn.com/problems/combinations/submissions/
//思路同括号问题
func combine(n int, k int) [][]int {
	var res [][]int
	_combine(1, n, k, []int{}, &res)
	return res
}

func _combine(min, max int, k int, ret []int, res *[][]int) {
	if len(ret) == k {
		//深拷贝
		cpret := make([]int, k, k)
		copy(cpret, ret)
		*res = append(*res, cpret)
		return
	}
	for min <= max {
		min++
		_combine(min, max, k, append(ret, min-1), res)
	}
}

func main() {

	//[[1,2,3,4],[1,2,3,5],[1,2,4,5],[1,3,4,5],[2,3,4,5]]
	fmt.Println(combine(5, 4))
}
