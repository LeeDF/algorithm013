package Week_03

import (
	"fmt"
	"sort"
)

//map 记录已经存在的
//优化， 先排序， 然后同一层， 如果当前元素等于前一个分支的元素，且前一个分支已经使用过，则剪枝
func permuteUnique(nums []int) [][]int {
	var result [][]int
	cach := make([]bool, len(nums))
	sort.Ints(nums)
	fmt.Println(nums)
	_permuteUnique(nums, len(nums), []int{}, &result, cach, )
	return result
}

func _permuteUnique(nums []int, max int, ret []int, result *[][]int, cach []bool) {
	if len(ret) == max {
		cpret := make([]int, max, max)
		copy(cpret, ret)
		*result = append(*result, cpret)
		return
	}

	for i := range nums{
		if !cach[i]{
			// 剪枝条件：i > 0 是为了保证 nums[i - 1] 有意义
			// 写 !used[i - 1] 是因为 nums[i - 1] 在深度优先遍历的过程中刚刚被撤销选择
			if i > 0 && nums[i] == nums[i-1] && !cach[i-1]{
				continue
			}
			cach[i] = true

			_permuteUnique(nums, max, append(ret, nums[i]), result, cach)
			cach[i] = false
		}
	}
}

func main() {
	fmt.Println(permuteUnique([]int{1, 1, 2}))
}
