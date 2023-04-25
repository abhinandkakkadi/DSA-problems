package main

import "fmt"

const size = 7
type HashTable struct {
	array [size]*bucket
}

type bucket struct {
	head *bucketNode
}

type bucketNode struct {
	key int
	next *bucketNode
}

func Init() *HashTable {
	result := &HashTable{}
	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return result
}

func main() {

	// find duplicates using hash table
	arr := []int{34,1,123,1,23,1,23,34}
	h := Init()
	for _,val := range arr {
		h.Insert(val)
	}

}


func (h *HashTable) Insert(value int) {

	index := hash(value)
	same := h.array[index].insert(value)
	if same {
		fmt.Print(value," ")
	}
}

func (b *bucket) insert(value int) bool {
	node := &bucketNode{value,nil}
	if b.head == nil {
		b.head = node
		return true
	}
	temp := b.head
	for temp != nil {
		if temp.key == value {
			return false
		}
		temp =  temp.next
	}
	
	node.next = b.head
	b.head = node
	return true

}	

func hash(value int) int {

	return value%size

}