package Test

import (
	"fmt"
	"math"
	"testing"
)

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
/** initialize your data structure here. */
type MinStack struct {
	stack []int
	minStack []int
}
func Constructor() MinStack {
	return MinStack{
		stack: []int{},
		minStack: []int{math.MaxInt64},
	}
}

func (this *MinStack) Min() int {
	return this.minStack[len(this.minStack)-1]
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (this *MinStack) Push(x int)  {
	this.stack = append(this.stack,x)
	top := this.minStack[len(this.minStack) - 1]
	this.minStack = append(this.minStack, min(x,top))
}

func (this *MinStack) Pop()  {
	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}
func Test(t *testing.T)  {
	stack := Constructor()
	stack.Push(5)
	stack.Push(3)
	stack.Push(2)
	stack.Push(1)
	stack.Push(-1)
	fmt.Println(stack.Min())
	stack.Pop()
	fmt.Println(stack.Min())
	fmt.Println(stack.Top())
	stack.Pop()
	fmt.Println(stack.Top() ,stack.Min())

}