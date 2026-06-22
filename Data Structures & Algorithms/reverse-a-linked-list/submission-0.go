/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
    result := &ListNode{}
	iter := result
	
	var traverse func(head *ListNode)
	traverse = func(head *ListNode) {
		if head == nil {
			return
		}
		traverse(head.Next)

		iter.Next = &ListNode{Val: head.Val}
		iter = iter.Next
	}

	traverse(head)
	
	return result.Next

}
