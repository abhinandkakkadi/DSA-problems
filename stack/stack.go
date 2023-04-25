package main

import "fmt"

func main() {

	arr := []int{54,23,34,23,2,45,34,23,45,23,4,34}

	selectionSort(arr)

	fmt.Println(arr)
}

func selectionSort(arr []int) {

	var small int
	var pos int

	for i :=0; i < len(arr); i++ {

		small = arr[i]
		pos = i

		for j:=i+1; j<len(arr); j++ {

			if arr[j] < small {

				small = arr[j]
				pos = j

			}
		}

		arr[i],arr[pos] = arr[pos],arr[i]
	}
}