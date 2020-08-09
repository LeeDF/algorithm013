package Week_02

//https://leetcode-cn.com/problems/binary-tree-preorder-traversal/
//根左右
func preorderTraversal(root *TreeNode) []int {
	ret := &[]int{}
	_preorderTraversal(root, ret)
	return *ret
}
func _preorderTraversal(root *TreeNode, ret *[]int) {
	if root != nil {
		*ret = append(*ret, root.Val)
		_preorderTraversal(root.Left, ret)
		_preorderTraversal(root.Right, ret)
	}
}
