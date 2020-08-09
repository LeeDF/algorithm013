package Week_02

//https://leetcode-cn.com/problems/n-ary-tree-level-order-traversal/

func levelOrder(root *Node) [][]int {
	var ret [][]int
	que := []*Node{root,}
	for len(que) > 0 {
		var tmp []int
		tq := que
		que = []*Node{}
		for _, i := range tq {
			if i != nil {
				tmp = append(tmp, i.Val)
				if len(i.Children) > 0 {
					que = append(que, i.Children...)
				}
			}

		}
		if len(tmp) > 0 {
			ret = append(ret, tmp)
		}

	}
	return ret
}
