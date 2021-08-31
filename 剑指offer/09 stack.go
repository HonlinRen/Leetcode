package main

import (
"container/list"
"fmt"
)

type CQueue struct {
	stack1, stack2 *list.List
}
func main(){
	que := Constructor()
	que.AppendTail(5)
	que.AppendTail(4)
	que.AppendTail(2)

	fmt.Println(que.DeleteHead())
	fmt.Println(que.DeleteHead())

	fmt.Println(que.DeleteHead())


}
func Constructor() CQueue {
	return CQueue{
		stack1: list.New(),
		stack2: list.New(),
	}
}

func (this *CQueue) AppendTail(value int)  {
	this.stack1.PushBack(value)
}

func (this *CQueue) DeleteHead() int {
	// 如果第二个栈为空
	if this.stack2.Len() == 0{
		for this.stack1.Len() >0{
			this.stack2.PushBack(this.stack1.Remove(this.stack1.Back()))
		}
	}
	if this.stack2.Len() != 0{
		e := this.stack2.Back()
		this.stack2.Remove(e)
		return e.Value.(int)
	}
	return -1
}
