package Week_03

import "fmt"

func subsets(nums []int) [][]int {
	var res [][]int
	_subsets(nums, 0, []int{}, &res)
	return res
}

func _subsets(nums []int, idx int, ret []int, res *[][]int) {
	cpret := make([]int, len(ret))
	copy(cpret, ret)
	*res = append(*res, cpret)
	for i:= idx;i < len(nums);i++ {
		_subsets(nums, i+1, append(ret, nums[i]), res)
	}

}

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}
