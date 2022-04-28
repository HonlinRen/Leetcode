package main

type node struct {
	data  int   // 数据域
	left  *node // 指向左子树的根结点
	right *node // 指向右子树的根结点
}

func NewNode(data int) *node {
	return &node{data: data}
}
