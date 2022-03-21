package main

import "fmt"

// 题目描述：
// 给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
//
//请你将两个数相加，并以相同形式返回一个表示和的链表。
//
//你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
//
// 输入：l1 = [2,4,3], l2 = [5,6,4]
// 输出：[7,0,8]
// 解释：342 + 465 = 807.
//
// 输入：l1 = [0], l2 = [0]
// 输出：[0]
// 输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
// 输出：[8,9,9,9,0,0,0,1]

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
	res := &ListNode{}
	carry := 0
	pre := res
	for l1 != nil || l2 != nil {
		// 当前list1 list2的取值
		n1, n2 := 0, 0
		// 如果l1和l2还有值，则取出来，并且更新节点位置
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}

		// 当前的和为 l1 l2的值加上前面的进位
		sum := n1 + n2 + carry
		// 当前存的值应该为 sum % 10
		val := sum % 10
		// 更新进位
		carry = sum / 10

		cur := &ListNode{val, nil}
		pre.Next = cur
		pre = pre.Next
	}

	if carry > 0 {
		pre.Next = &ListNode{carry, nil}
	}

	return res.Next
}
