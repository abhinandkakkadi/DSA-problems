package main

import "fmt"

type Node struct {
	left *Node
	key int
	right *Node
}

type BinarySearchTree struct {
	root *Node
}

func main() {

	s := &BinarySearchTree{}
	s.Insert(5)
	s.Insert(4)
	s.Insert(3)

	Inoder(s.root)

	se := Search(s.root,4)
	fmt.Println(se)

	s.Delete(4)
	fmt.Println(se)
}

func (s *BinarySearchTree) Insert(value int) {

	s.root = addValue(s.root,value)

}

func addValue(root *Node, value int) *Node {

	if root == nil {
		node := &Node{nil,value,nil}
		return node
	}

	if value > root.key {
		root.right = addValue(root.right,value)
	} else {
		root.left = addValue(root.left,value)
	}

	return root
}

func Inoder(root *Node) {

	if root == nil {
		return
	}

	Inoder(root.left)
	fmt.Print(root.key," ")
	Inoder(root.right)
}

func Search(root *Node, value int) bool {

	if root == nil {
		return false
	}

	if root.key == value {
		return true
	}

	if value > root.key {
		return Search(root.right,value)
	} else {
		return Search(root.left,value)
	}
}

func (s *BinarySearchTree) Delete(value int) {

	s.root = delete(s.root,value)
}

func delete(root *Node,value int) *Node {

	if root == nil {
		return nil
	}

	if value > root.key {
		root.right = delete(root.right,value)
	} else if value < root.key {
		root.left = delete(root.left,value)
	} else {
		if root.left == nil {
			return root.right
		} else if root.right == nil {
			return root.left
		} else {
			c := success(root)
			root.key = c.key
			root.right = delete(root.right,c.key)
		}
	}

	return root
}

func success(root *Node) *Node {
	c := root.right

	for c != nil && c.left != nil {
		c = c.left
	}
	return c
}