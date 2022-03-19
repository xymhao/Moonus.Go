package treenode

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	first := head
	second := head.Next

	for second != nil && second.Next != nil {
		second = second.Next.Next
		first = first.Next
	}
	if second == nil {
		return first
	}

	return first.Next
}

//https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/
// 1-2-3-4-5 n:2
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head.Next == nil {
		return nil
	}
	// node1 先 next n 次
	node1 := head
	node2 := head

	for i := 0; i < n; i++ {
		node1 = node1.Next
	}

	if node1 == nil {
		return head.Next
	}

	for ; node1.Next != nil; node1 = node1.Next {
		node2 = node2.Next
	}
	node2.Next = node2.Next.Next
	return head
}
