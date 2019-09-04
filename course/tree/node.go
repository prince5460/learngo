package tree

import "fmt"

//结构体和方法

//值接收者和指针接收者:
//要改变内容必须只有指针接收者
//结构过大也考虑使用指针接收者
//一致性:如有指针接收者,最好都是指针接收者

type Node struct {
	Value       int
	Left, Right *Node
}

//自定义工厂函数
func CreateTreeNode(value int) *Node {
	return &Node{Value: value}
}

//为结构定义方法
func (node Node) PrintValue() {
	fmt.Println(node.Value)
}

//使用指针作为方法接收者,只有使用指针才可以改变结构内容
func (node *Node) SetValue(value int) {
	node.Value = value
}
