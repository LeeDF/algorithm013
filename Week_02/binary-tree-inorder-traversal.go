package Week_02

// Definition for a binary tree node.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//https://leetcode-cn.com/problems/binary-tree-inorder-traversal/submissions/
//左根右
func inorderTraversal(root *TreeNode) []int {
	ret := &[]int{}
	_inorderTraversal(root, ret)
	return *ret
}
func _inorderTraversal(root *TreeNode, ret *[]int) {
	if root != nil {
		_inorderTraversal(root.Left, ret)
		*ret = append(*ret, root.Val)
		_inorderTraversal(root.Right, ret)
	}
}
