package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_ConstructNode(t *testing.T) {
	listNode := ListNode{}
	listNode.ConstructListNode([]int{1, 2, 3})
}

func Test_PrintListNode(t *testing.T) {
	listNode := ListNode{}
	listNode.ConstructListNode([]int{1, 2, 3})
	str := listNode.String()
	fmt.Println(str)
}

func Test_Should_Return807(t *testing.T) {
	listNode1 := &ListNode{}
	listNode1.ConstructListNode([]int{2, 4, 3})
	listNode2 := &ListNode{}
	listNode2.ConstructListNode([]int{5, 6, 4})
	expect := &ListNode{
		Val: 7,
		Next: &ListNode{
			0,
			&ListNode{
				8,
				nil,
			},
		},
	}
	res := addTwoNumbers(listNode1, listNode2)
	fmt.Println(res)
	fmt.Println(expect)
	if !reflect.DeepEqual(res, expect) {
		t.Error("Should be correct")
	}

}
