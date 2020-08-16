package Week_03

import (
	"math"
)
//ret的长度就是行数， 遍历列即可
func solveNQueens(n int) [][]string {
	var res [][]string
	_solveNQueens(n, 0, []int{}, &res)
	return res
}

func _solveNQueens(n int, col int, ret []int, res *[][]string) {
	if len(ret) == n {
		*res = append(*res, _String(ret))
		return
	}

	for i := 0; i < n; i++ {
		if !_inattack(ret, i) {
			ret = append(ret, i)
			_solveNQueens(n, i, ret, res)
			ret = ret[:len(ret)-1]
		}
	}
}

func _String(ret []int) []string {
	res := make([]string, len(ret))
	for k, i := range ret {

		for j := 0; j < len(ret); j++ {
			s := "."
			if j == i{
				s = "Q"
			}
			res[k] += s
		}
	}
	return res
}

//判断攻击范围
func _inattack(ret []int, col int) bool {
	for i := range ret {
		//同列攻击
		if col == ret[i] {
			return true
		}
		//左斜线和右斜线的斜率分别是-1 和 1， 斜率 = (y1-y0)/ (x1-x0)
		t := float64(col-ret[i]) / float64(len(ret)-i)
		if math.Abs(t) == 1 {
			return true
		}
	}
	return false
}

func main() {
	solveNQueens(4)
}
