package main

import (
	"fmt"
	"learngo/course/tree"
)

//扩展已有类型

//组合
type myTreeNode struct {
	node *tree.Node
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}

	left := myTreeNode{myNode.node.Left}
	right := myTreeNode{myNode.node.Right}

	left.postOrder()
	right.postOrder()
	myNode.node.PrintValue()

}

func main() {
	var root tree.Node

	//treeNode的创建方法
	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)

	nodes := []tree.Node{
		{Value: 3},
		{},
		{6, nil, &root},
	}
	fmt.Println(nodes)

	root.Left.Right = tree.CreateTreeNode(2)

	root.Right.Left.SetValue(4)
	root.Right.Left.PrintValue()

	//root.printValue()
	//root.setValue(100)
	//root.printValue()

	fmt.Println("------traverse()------")

	root.Traverse()

	fmt.Println("--------------")
	myRoot := myTreeNode{&root}
	myRoot.postOrder()

}
