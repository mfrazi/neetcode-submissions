/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeElements(head *ListNode, val int) *ListNode {
	result := &ListNode{Next: head}
	iter := result

	for iter != nil {
		if iter.Next == nil || iter.Next.Val != val {
			iter = iter.Next
			continue
		}
		iter.Next = iter.Next.Next
	}
	return result.Next
}
