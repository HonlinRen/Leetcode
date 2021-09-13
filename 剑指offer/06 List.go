package main

import (
	"fmt"
)
// 尾部开始 输出链表

func (list ListNode) prinfListInReserveOrder() {
	if list.Next != nil {
		list.Next.prinfListInReserveOrder()
	}
	fmt.Print("<-  ", list.Val)
}
func (list ListNode) prinfListInOrder() {
	fmt.Print("  ->", list.Val)
	if list.Next != nil {
		list.Next.prinfListInOrder()
	}
}
func (list *ListNode) Push(v int) {
	TempList := list
	//走到尾部
	for TempList.Next != nil {
		TempList = TempList.Next
	}
	TempList.Next = &ListNode{
		Val: v,
	}
}
func (list *ListNode) Set(v int) {
	list.Val = v
}


func main(){
	head := &ListNode{}
	for i := 0; i < 3; i++ {
		head.Push(i)
	}
	head.prinfListInReserveOrder()
}
