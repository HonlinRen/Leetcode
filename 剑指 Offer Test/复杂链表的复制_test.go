package Test

import (
	"fmt"
	"testing"
)

const value = 0
const random = 1
func TestName(t *testing.T) {
	head := [5][2]int{{7,-1},{13,0},{11,4},{10,2},{1,0}}
	var createNode map[int]*Node
	createNode = make(map[int]*Node)


	for i := 0; i < len(head); i++ {
		createNode[i] = &Node{}
		createNode[i].Val = head[i][value]
		if i > 0{
			createNode[i - 1].Next = createNode[i]
		}

	}
	for i := 0; i < len(head); i++ {
		if head[i][random] == -1{
			createNode[i].Random = nil
		}else {
			createNode[i].Random = createNode[head[i][random]]
		}
	}
	nodeCopy := copyRandomList(createNode[0])
	for nodeCopy != nil{
		fmt.Println(nodeCopy.Val)
		nodeCopy = nodeCopy.Next
	}

}
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}
func (head *Node) push(x int) *Node{
	head.Val = x
	return head.Next
}
// test
var cachedNode map[*Node]*Node

func deepCopy(node *Node) *Node {
	if node == nil {
		return nil
	}
	if n, has := cachedNode[node]; has {
		return n
	}
	newNode := &Node{Val: node.Val}
	cachedNode[node] = newNode
	newNode.Next = deepCopy(node.Next)
	newNode.Random = deepCopy(node.Random)
	return newNode
}

func copyRandomList(head *Node) *Node {
	cachedNode = map[*Node]*Node{}
	return deepCopy(head)
}



