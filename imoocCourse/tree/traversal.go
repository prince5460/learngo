package tree

//遍历
func (node *Node) Traverse() {
	if node == nil {
		return
	}
	node.Left.Traverse()
	node.PrintValue()
	node.Right.Traverse()
}
