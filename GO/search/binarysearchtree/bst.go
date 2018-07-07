package binarysearchtree

type Node struct {
	key int
	value int
	left *Node
	right *Node
}

type BST struct {
	root *Node
	count int
}


func (r BST) Construct() {

	r.count = 0
}


func (r BST) Size() int{
	return r.count
}


func (r BST) IsEmpty() bool {
	return r.count == 0
}