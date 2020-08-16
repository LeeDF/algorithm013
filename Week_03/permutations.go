package Week_03

import "fmt"

//https://leetcode-cn.com/problems/permutations/

func permute(nums []int) [][]int {
	var res [][]int
	cach := make([]bool, len(nums))
	_permute(nums, []int{}, len(nums), &res, cach)
	return res
}

func _permute(nums []int, ret []int, max int, res *[][]int, cach []bool) {
	if len(ret) == max {
		//深拷贝
		cpret := make([]int, max, max)
		copy(cpret, ret)
		*res = append(*res, cpret)
		return
	}
	for i := 0; i < len(nums); i++ {
		if !cach[i] {
			cach[i] = true
			_permute(nums, append(ret, nums[i]), max, res, cach)
			cach[i] = false
		}

	}

}

func main() {
	fmt.Println(permute([]int{1, 1, 2}))
}
