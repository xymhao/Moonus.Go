package treenode

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNode(t *testing.T) {
	node := middleNode(&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}})

	assert.Equal(t, 3, node.Val)

	node2 := middleNode(&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}})

	assert.Equal(t, 3, node2.Val)

	node3 := middleNode(&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6}}}}}})

	assert.Equal(t, 4, node3.Val)

}

func TestRemoveNthFromEnd(t *testing.T) {
	removeNthFromEnd(&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: nil}}, 2)

	removeNthFromEnd(&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4,
		Next: &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 7, Next: &ListNode{Val: 8, Next: &ListNode{Val: 9,
			Next: &ListNode{Val: 10}}}}}}}}}}, 7)

	removeNthFromEnd(&ListNode{Val: 1, Next: &ListNode{Val: 2}}, 1)

	//removeNthFromEnd(&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}, 2)

	end := removeNthFromEnd(&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}}, 2)
	fmt.Println(end)
	removeNthFromEnd(&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6}}}}}}, 2)
}
