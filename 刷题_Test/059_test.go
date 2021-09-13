package 刷题_Test

import (
	"fmt"
	"testing"
)


func Test059(t *testing.T) {
	constructor := Constructor()

	constructor.Push_back(5)
	constructor.Push_back(6)
	constructor.Push_back(6)
	constructor.Push_back(2)
	constructor.Push_back(4)

	fmt.Println(constructor.q1,constructor.max)
}

type MaxQueue struct {
	q1 []int
	max []int
}

func Constructor() MaxQueue {
	return MaxQueue{
		make([]int,0),
		make([]int,0),
	}
}

func (this *MaxQueue) Max_value() int {
	if len(this.max) == 0{
		return -1
	}
	return this.max[0]
}
//入队 两个必入
func (this *MaxQueue) Push_back(value int)  {
	this.q1 = append(this.q1,value)
	for len(this.max) > 0 && value > this.max[len(this.max) - 1]{
		this.max = this.max[:len(this.max) - 1]
	}
	this.max = append(this.max,value)
}
//出队 有可能max不出
func (this *MaxQueue) Pop_front() int {
	//队列为空 返回-1
	n := -1
	if len(this.q1) > 0{
		n = this.q1[0]
		this.q1 = this.q1[1:]
		if n == this.max[0]{
			this.max = this.max[1:]
		}
	}

	return n
}

