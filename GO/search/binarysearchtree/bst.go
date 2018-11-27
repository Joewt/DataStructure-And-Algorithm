package main

import (
	"fmt"
	"math/rand"
	"time"
)

//节点

type Key int
type Value int

type Node struct {
	key   Key
	value Value
	left  *Node
	right *Node
}

//搜索树定义
type BST struct {
	root  *Node
	count int
}

func (r *BST) Construct() {
	r.root = nil
	r.count = 0
}

func (r BST) Size() int {
	return r.count
}

func (r BST) IsEmpty() bool {
	return r.count == 0
}

func (r *BST) Insert(key Key, value Value) {
	r.root = r.insert(r.root, key, value)
}

func (r *BST) Contain(key Key) bool {
	return r.contain(r.root, key)
}

func (r *BST) Search(key Key) *Value {
	return r.search(r.root, key)
}

func (r *BST) PreOrder() {
	r.preOrder(r.root)
}

func (r *BST) InOrder() {
	r.inOrder(r.root)
}

func (r *BST) PostOrder() {
	r.postOrder(r.root)
}

func (r *BST) insert(node *Node, key Key, value Value) *Node {

	if node == nil {
		r.count++
		return &Node{key: key, value: value, left: nil, right: nil}
	}

	if key == node.key {
		node.value = value
	} else if key < node.key {
		node.left = r.insert(node.left, key, value)
	} else {
		node.right = r.insert(node.right, key, value)
	}

	return node
}

func (r *BST) contain(node *Node, key Key) bool {
	if node == nil {
		return false
	}

	if key == node.key {
		return true
	} else if key < node.key {
		return r.contain(node.left, key)
	} else {
		return r.contain(node.right, key)
	}
}

func (r *BST) search(node *Node, key Key) *Value {
	if node == nil {
		return nil
	}

	if key == node.key {
		return &node.value
	} else if key < node.key {
		return r.search(node.left, key)
	} else {
		return r.search(node.right, key)
	}
}

func (r *BST) preOrder(node *Node) {
	if node != nil {
		fmt.Println(node.key)
		r.preOrder(node.left)
		r.preOrder(node.right)
	}
}

func (r *BST) inOrder(node *Node) {
	if node != nil {
		fmt.Println(node.key)
		r.inOrder(node.left)
		r.inOrder(node.right)
	}
}

func (r *BST) postOrder(node *Node) {
	if node != nil {
		fmt.Println(node.key)
		r.postOrder(node.left)
		r.postOrder(node.right)
	}
}

func main() {
	bst := BST{root: nil, count: 0}
	const N = 100
	const M = 100
	for i := 0; i < N; i++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		k := Key(r.Intn(M))
		v := Value(k)
		fmt.Printf("%d ", k)
		bst.Insert(k, v)
	}

	fmt.Printf("size %d\n", bst.Size())

	bst.InOrder()
}
