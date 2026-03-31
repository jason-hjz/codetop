package _06_reverseList

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

func initList(val []int) *ListNode {
	if len(val) == 0 {
		return nil
	}
	head := &ListNode{Val: val[0]}
	curr := head
	for i := 1; i < len(val); i++ {
		curr.Next = &ListNode{Val: val[i]}
		curr = curr.Next
	}
	return head
}

func traverseList(head *ListNode) {
	curr := head
	for curr != nil {
		fmt.Println(curr.Val)
		curr = curr.Next
	}
}
