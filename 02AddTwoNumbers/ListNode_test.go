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

func Test_ShouldBeCorrect(t *testing.T) {
	listNode1 := &ListNode{}
	listNode1.ConstructListNode([]int{9, 9, 9, 9, 9, 9, 9})

	listNode2 := &ListNode{}
	listNode2.ConstructListNode([]int{9, 9, 9, 9})

	expect := &ListNode{
		Val: 8,
		Next: &ListNode{
			9,
			&ListNode{
				9,
				&ListNode{
					0,
					&ListNode{
						1,
						nil,
					},
				},
			},
		},
	}

	res := addTwoNumbers(listNode1, listNode2)
	fmt.Println(listNode1)
	fmt.Println(listNode2)
	fmt.Println(expect)
	if !reflect.DeepEqual(res, expect) {
		t.Error("Should be correct")
	}
}
