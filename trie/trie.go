package main

import "fmt"


func main() {

	arr := []int{24,23,345,23,456,7}
	for i:= (len(arr)/2)-1; i>=0; i-- {
		Heapify(arr,i,len(arr)-1)
	}
	fmt.Println(arr)
}

func leftChild(i int) int {
	return 2*i +1
}

func rightChild(i int) int {
	return 2*i + 2
}

func Heapify(arr []int,index int, lastIndex int) {

	compareChild := index
	l,r := leftChild(index),rightChild(index)

	for l <= lastIndex {
		if l == lastIndex {
			compareChild = l
		} else if arr[l] > arr[r] {
			compareChild = l
		} else {
			compareChild =r
		}

		if arr[index] < arr[compareChild] {
			arr[index],arr[compareChild] = arr[compareChild],arr[index]
			index = compareChild
			l,r = leftChild(index),rightChild(index)
		} else {
			return
		}
	}
}