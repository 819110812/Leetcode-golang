package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	var s string
	for l != nil {
		s += fmt.Sprintf("%d", l.Val)
		l = l.Next
	}
	return s
}

func (l *ListNode) ConstructListNode(nums []int) {
	l.Val = nums[0]
	temp := l
	for i := 1; i < len(nums); i++ {
		temp.Next = &ListNode{Val: nums[i]}
		temp = temp.Next
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	pre := &ListNode{0, nil}
	temp := 0
	res := pre
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := n1 + n2 + temp
		temp = sum / 10
		sum = sum % 10
		curList := &ListNode{sum, nil}
		res.Next = curList
		res = res.Next
		//if l1 != nil {
		//	l1 = l1.Next
		//}
		//if l2 != nil {
		//	l2 = l2.Next
		//}

	}
	if temp > 0 {
		res.Next = &ListNode{temp, nil}
	}
	return pre.Next
}
