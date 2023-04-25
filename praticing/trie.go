package main

import "fmt"

const size = 26
type Node struct {
	children [size]*Node
	isEnd bool
}

type Trie struct {
	root *Node
}

func (t *Trie) Insert(w string) {
	
	currentNode := t.root

	for i:=0; i< len(w); i++ {
		currentLetter := w[i] - 'a'
		if currentNode.children[currentLetter] == nil {
			currentNode.children[currentLetter] = &Node{}
		}
		currentNode = currentNode.children[currentLetter]
	}
	currentNode.isEnd = true
	
}

func (t *Trie) Search(w string) bool {

	currentNode := t.root

	for i:=0; i < len(w); i++ {

		currentLetter := w[i] - 'a'
		if currentNode.children[currentLetter] == nil {
			return false
		}
		currentNode = currentNode.children[currentLetter]

	}
	if currentNode.isEnd == true {
		return true
	} else {
		return false
	}
}

func InitTrie() *Trie {
	result := &Trie{root:&Node{}}
	return result
}



func main() {

	t := InitTrie()

	t.Insert("abhinand")

	fmt.Println(t.Search("abhinand"))
}