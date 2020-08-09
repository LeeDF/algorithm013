package Week_02

//https://leetcode-cn.com/problems/n-ary-tree-postorder-traversal/

//N叉树后序遍历
//Definition for a Node.
type Node struct {
	Val      int
	Children []*Node
}

//左右中
func postorder(root *Node) []int {
	ret := &[]int{}
	_postorder(root, ret)
	return *ret
}

func _postorder(root *Node, ret *[]int) {
	if root != nil {
		for _, i := range root.Children {
			_postorder(i, ret)
		}
		*ret = append(*ret, root.Val)
	}
}
