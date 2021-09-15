package 刷题_Test

import (
	"testing"
)

func Test007(t *testing.T){

}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0{
		return nil
	}
	root := &TreeNode{preorder[0],nil,nil}
	i:=0
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}
	//left 只取前半段
	root.Left = buildTree(preorder[1:i+1],inorder[:i])
	//right 只取后半段
	root.Right = buildTree(preorder[i+1:],inorder[i+1:])
	return root
}