package main

import "fmt"

type treeNode struct {
	value int
	left, right *treeNode
}

func (node treeNode) print()  {
	fmt.Println(node.value, " ")
}

func (node *treeNode) setValue(value int)  {
	node.value = value
}
//遍历
func (node *treeNode) traverse()  {
	if node == nil{
		return
	}

	node.left.traverse()
	node.print()
	node.right.traverse()
}

func createNode(value int) *treeNode {
	return &treeNode{value:value}
}



func main()  {
	root := treeNode{value:3}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil,nil}

	root.traverse()
}