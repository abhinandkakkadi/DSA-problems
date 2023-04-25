package main

import (
	"fmt"
	"math"
)

type Node struct {
	left  *Node
	key   int
	right *Node
}

type binarySearchTree struct {
	root *Node
}

func main() {

	b := binarySearchTree{}
	b.Insert(2)
	b.Insert(3)
	b.Insert(4)
	b.Insert(5)
	b.Insert(5)

	inOrder(b.root)
	fmt.Println()

	f := Search(b.root,3)
	fmt.Println(f)
	fmt.Println()

	b.Delete(3)
	b.Delete(2)
	f = Search(b.root,3)
	fmt.Println(f)
	inOrder(b.root)
	fmt.Println()

	tr := isBST(b.root,math.MinInt,math.MaxInt)
	fmt.Println(tr)

}

func (b *binarySearchTree) Insert(value int) {

	b.root = insert(b.root, value)
}

func insert(root *Node, value int) *Node {

	if root == nil {

		node := &Node{nil, value, nil}
		return node

	}

	if root.key == value {
		return root
	}

	if value > root.key {
		root.right = insert(root.right, value)
	} else {
		root.right = insert(root.left, value)
	}

	return root
}

func inOrder(root *Node) {

	if root == nil {
		return
	}

	inOrder(root.left)
	fmt.Print(root.key," ")
	inOrder(root.right)
}


func Search(root *Node,value int) bool {

	if root == nil {
		return false
	}

	if value > root.key {
		return Search(root.right,value)
	} else if value < root.key {
		return Search(root.left,value)
	} else {
		return true
	}
}


func (b *binarySearchTree) Delete(value int) {

	b.root = delete(b.root,value)
}

func delete(root *Node, value int) *Node {

	if root == nil {
		return nil
	}

	if value < root.key {
		root.left = delete(root.left, value)
	} else if value > root.key {
		root.right = delete(root.right, value)
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


func success(c *Node) *Node {

	c = c.right 

	for c != nil && c.left != nil {
		c = c.left
	}

	return c
}

func isBST(root *Node, min int, max int ) bool {

	if root == nil {
		return true
	}

	return root.key > min && root.key < max && isBST(root.left,min,root.key) && isBST(root.right,root.key,max)
}