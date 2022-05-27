package problems

import (
	"fmt"
	"strings"
)

// You are given the heads of two sorted linked lists list1 and list2.

// Merge the two lists in a one sorted list. The list should be made by splicing together the nodes of the first two lists.

// Return the head of the merged linked list.


type ListNode struct {
	Val int
	Next *ListNode
}

func (n *ListNode) PrintList () string {
	
	var sb strings.Builder
	next := n
	for next != nil {
		sb.WriteString(fmt.Sprintf("%d ", n.Val))
		next = next.Next
	}
	return sb.String()
}


 func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	var head *ListNode // we can remove this and just use list1 and list2 in the end
	var current *ListNode

	p1 := list1
	p2 := list2

	if p1 == nil {
		return p2
	}

	if p2 == nil {
		return p1
	}
	
	if (p1.Val > p2.Val) {
		head = p2
		p2 = p2.Next
	} else {
		head = p1
		p1 = p1.Next
	}
	current = head//  this point to the current
	current.Next = nil //this is not needed just to denote this is independent now and we are justs building new list

	for p1 != nil && p2 != nil {

		if (p1.Val <= p2.Val) {
			current.Next = p1
			current = p1
			p1 = p1.Next
		} else {
			current.Next = p2
			current = p2
			p2 = p2.Next

		}
	}

	if p1 == nil {
		current.Next = p2
	} else {
		current.Next = p1
	}

    return head
}