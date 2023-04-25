package main

import (
	"fmt"
	"math"
)

type Node struct {
	left *Node
	data int
	right *Node
}

type BinarySearchTree struct {
	root *Node
}

func (b *BinarySearchTree) Insert(value int) {
	b.root = addNode(b.root,value)
}

func  addNode(root *Node, value int ) *Node {

	if root == nil {
		node := &Node{nil,value,nil}
		return node
	}

	if root.data == value {
		return nil
	}

	if value > root.data {
		root.right = addNode(root.right,value)
	} else {
		root.left = addNode(root.left,value)
	}

	return root
}

func Search(root *Node, value int) bool {

	if root == nil {
		return false
	}

	

	if value > root.data {
		return Search(root.right,value)
	} else if value < root.data {
		return Search(root.left,value)
	} else {
		return true
	}
}

func (b *BinarySearchTree) Delete(value int) {
	b.root = delete(b.root,value)
}

func delete(root *Node, value int) *Node {

	if root == nil {
		return nil
	}

	if value > root.data {
		root.right = delete(root.right,value)
	} else if value < root.data {
		root.left = delete(root.left,value)
	} else {
		if root.left == nil {
			return root.right
		} else if root.right == nil {
			return root.left
		} else {
			c := success(root)
			root.data = c.data
			root.right = delete(root.right,c.data)
		}
	}
	return root
}

func success(c *Node) *Node {
	c = c.right
	for c != nil && c.left != nil {
		c = c.left
	}
	return c
}




func main() {

	b := BinarySearchTree{}
	b.Insert(3)
	b.Insert(10)
	b.Insert(4)
	b.Insert(23)
	inOrder(b.root)
	
	// fmt.Println(Search(b.root,2))
	// b.Delete(2)
	// b.Delete(3)
	// fmt.Println(Search(b.root,2))
	// inOrder(b.root)

	fmt.Println(isBST(b.root,math.MinInt,math.MaxInt))

	fmt.Println(closest(b.root,2))
}

func inOrder(root *Node) {

	if root == nil {
		return
	}

	inOrder(root.left)
	fmt.Println(root.data)
	inOrder(root.right)
	
}


func isBST(root *Node, min int, max int) bool {

	if root == nil {
		return true
	}

	return root.data > min && root.data < max && isBST(root.left,min,root.data) && isBST(root.right,root.data,max)
}


func closest(root *Node, value int) int {

	if root == nil {
		return -1
	}
	minDiff := math.MaxInt
	var closest int

	for root != nil {
		currentDiff := math.Abs(float64(root.data)-float64(value))

		if int(currentDiff) < minDiff {
			minDiff = int(currentDiff)
			closest = root.data
		}

		if value > root.data {
			root = root.right
		} else if  value < root.data {
			root = root.left
		} else {
			return root.data
		}
	}

	return closest
	
}