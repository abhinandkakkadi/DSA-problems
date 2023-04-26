package main

import (
	"dsa/algorithms/algorithm"
	"fmt"
)


func main() {

	// selection sort
	arr1 := []int{23,1,34,23,34,23,2}
	algorithm.SelectionSort(arr1,len(arr1))
	fmt.Println(arr1)

	// Buble sort
	arr2 := []int{12,1,23,4,1,34}
	algorithm.BubbleSort(arr2,len(arr2))
	fmt.Println(arr2)

	// insertion sort
	arr3 := []int{343,2,12,34,2,3,23}
	algorithm.InsertionSort(arr3,len(arr3))
	fmt.Println(arr3)

	// quick sort
	arr4 := []int{3453,2,23,1,11,123,2,45}
	algorithm.QuickSort(arr4,0,len(arr4)-1)
	fmt.Println(arr4)

}


