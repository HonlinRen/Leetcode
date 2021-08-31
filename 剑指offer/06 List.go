package main
import (
	"fmt"
	"strconv"
)
// 尾部开始 输出链表

type LinkedList struct {
	Value string
	Next  *LinkedList
}
func (list LinkedList) prinfListInReserveOrder() {
	if list.Next != nil {
		list.Next.prinfListInReserveOrder()
	}
	fmt.Print("<-  ", list.Value)
}
func (list LinkedList) prinfListInOrder() {
	fmt.Print("  ->", list.Value)
	if list.Next != nil {
		list.Next.prinfListInOrder()
	}
}
func (list *LinkedList) Push(v string) {
	TempList := list
	//走到尾部
	for TempList.Next != nil {
		TempList = TempList.Next
	}
	TempList.Next = &LinkedList{
		Value: v,
	}
}
func (list *LinkedList) Set(v string) {
	list.Value = v
}
func sword3() {
	list := &LinkedList{
		Value: "head",
	}
	for i := 0; i < 10; i++ {
		list.Push(strconv.Itoa(i))
	}
	fmt.Println("PinrfInOrder")
	list.prinfListInOrder()
	fmt.Println("")
	fmt.Println("PinrfInReservedOrder")
	list.prinfListInReserveOrder()
}
func main(){
	sword3()
}
