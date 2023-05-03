package main

import (
	"dsa/basicarray/array"
	"fmt"
)

func main() {

	// largest element in an array
	arr1 := []int{343,23,23,3,4,23,34,2,2323}
	fmt.Println(array.LargestElement(arr1))	

	// second largest element in an array
	arr2 := []int{23,2123,12,545,23,2,112}
	fmt.Println(array.SecondLargest(arr2))

	// Remove duplicates from sorted array
	arr3 := []int{1,2,2,2,3,3,3,4,4,5}
	size := array.RemoveDuplicatesFromSortedArray(arr3)
	for i :=0; i < size; i++ {
		fmt.Print(arr3[i])
	}
	fmt.Println()

	// left rotate array by one
	arr4 := []int{1,2,3,4,5}
	array.LeftRotate(arr4,2)
	fmt.Println(arr4)

	// move zeroes to the end

	arr5 := []int{0,23,0,2,0,0,25}
	array.MoveZeroes(arr5)
	fmt.Println(arr5)

}