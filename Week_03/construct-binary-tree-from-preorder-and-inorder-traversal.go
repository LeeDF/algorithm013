package Week_03

import "fmt"

//https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/solution/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//[ 根节点, [左子树的前序遍历结果], [右子树的前序遍历结果] ]
//[ [左子树的中序遍历结果], 根节点, [右子树的中序遍历结果] ]
//且len[左子树的中序遍历结果] == len[左子树的前序遍历结果]
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	node := &TreeNode{Val: preorder[0]}
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}
	node.Left = buildTree(preorder[1:i+1],inorder[:i])
	node.Right = buildTree(preorder[i+1:], inorder[i+1:])
	return node
}

func main()  {
	a := []int{0,1,2,3,4}
	fmt.Println(a[1:2])
}
