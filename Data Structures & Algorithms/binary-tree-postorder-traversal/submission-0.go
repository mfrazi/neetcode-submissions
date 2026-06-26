/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func postorderTraversal(root *TreeNode) []int {
    result := []int{}
	dfs(root, &result)
	return result
}

func dfs(node *TreeNode, result *[]int) {
	if node == nil {
		return
	}

	dfs(node.Left, result)
	dfs(node.Right, result)
	*result = append(*result, node.Val)
}