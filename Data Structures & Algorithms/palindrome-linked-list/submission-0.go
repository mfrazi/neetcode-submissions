/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func isPalindrome(head *ListNode) bool {
	tail := head
	result := true
	var traverse func(tail *ListNode)
	traverse = func(tail *ListNode) {
		if tail == nil {
			return
		}
		traverse(tail.Next)
		if tail.Val != head.Val || !result {
			result = false
			return
		}
		head = head.Next
	}
	traverse(tail)

	return result
}
