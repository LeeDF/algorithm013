package Week_03

//pq在root的两侧， 则root就是公共祖先
//左右寻找pq， 找到就终止递归
//特殊情况，pq在root的同一侧
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root.Val == p.Val || root.Val == q.Val {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil{
		return root
	}
	if left != nil && right == nil{
		return left
	}
	if left == nil && right != nil{
		return right
	}
	return nil
}