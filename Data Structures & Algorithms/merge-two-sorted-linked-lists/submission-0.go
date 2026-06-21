/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    root := &ListNode{}
	iter := root

	for list1 != nil || list2 != nil {	
		if list1 != nil && list2 != nil {
			if list1.Val <= list2.Val {
				iter.Next = &ListNode{Val: list1.Val}
				list1 = list1.Next
			} else {
				iter.Next = &ListNode{Val: list2.Val}
				list2 = list2.Next
			}
		} else if list1 != nil {
			iter.Next = &ListNode{Val: list1.Val}
			list1 = list1.Next
		} else {
			iter.Next = &ListNode{Val: list2.Val}
			list2 = list2.Next
		}

		iter = iter.Next
	}

	return root.Next
}
