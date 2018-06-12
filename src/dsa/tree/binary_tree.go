package tree

type BinaryTree struct {
	left   BinaryTree
	right  BinaryTree
	parent BinaryTree
	value  int
}

func NewBinaryTree(v int) (*BinaryTree) {
	bt := new(BinaryTree)
	bt.parent = nil
	bt.value = v
	return bt
}
func (bt *BinaryTree) Root() int{
	p := bt.parent
	for p != nil {
		p=bt.parent
	}
	return p.value

}

//init
//Clear
//Root
//Value
//Assign
//Parent
//LeftChild
//RightChild
//LeftSibling
//RigttSibling
//InsertChild
//DeleteChild
//PreOrderTraverse
//InOrderTraverse
//PostOrderTraverse
//LevelOrderTraverse
