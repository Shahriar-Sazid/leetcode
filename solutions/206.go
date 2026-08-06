package solutions

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}

	array := []int{}

	cur := head
	for {
		array = append(array, cur.Val)
		if cur.Next == nil {
			break
		} else {
			cur = cur.Next
		}
	}

	var revHead *ListNode = &ListNode{}
	cur = revHead
	for i := len(array) - 1; i >= 0; i-- {
		cur.Val = array[i]
		if i > 0 {
			cur.Next = &ListNode{}
			cur = cur.Next
		} else {
			cur.Next = nil
		}
	}

	return revHead
}
