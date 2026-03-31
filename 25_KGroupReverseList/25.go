package _5_KGroupReverseList

type ListNode struct {
	Val  int
	Next *ListNode
}

// K个一组翻转链表、组内翻转

func reverseKGroup(head *ListNode, k int) *ListNode {
	// 1. 创建「虚拟头节点 hair」（哨兵节点），Next 指向原始头节点 head
	// 作用：统一处理链表头部的翻转（避免单独处理头节点的边界问题）
	hair := &ListNode{Next: head}

	pre := hair
	for head != nil {
		tail := pre
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return hair.Next
			}
		}
		next := tail.Next                    //当前组尾的下一个节点
		head, tail = reverseList(head, tail) //组内翻转返回新头尾
		pre.Next = head
		tail.Next = next
		pre = tail
		head = next
	}

	return hair.Next
}

func reverseList(head, tail *ListNode) (*ListNode, *ListNode) {
	prev := tail.Next
	p := head
	for prev != tail {
		nex := p.Next
		p.Next = prev
		prev = p
		p = nex
	}
	return tail, head
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
