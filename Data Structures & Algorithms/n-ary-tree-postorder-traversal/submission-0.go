/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func postorder(root *Node) []int {
	result := []int{}
	dfs(root, &result)
	return result
}

func dfs(node *Node, result *[]int) {
	if node == nil {
		return
	}

	for i:=0; i<len(node.Children); i++ {
		dfs(node.Children[i], result)
	}
	*result = append(*result, node.Val)
}