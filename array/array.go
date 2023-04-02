package main

import (
	"fmt"
	"math"
)


func main() {

	// second smallest element in an array
	arr := []int{23,1,34,34,4}
	var smallest int = math.MaxInt
	var sSmallest int = math.MaxInt

	s,ss := second_smallest(arr,smallest,sSmallest)
	fmt.Println("The smallest and second smallest are ",s," ",ss)


	// Reverse an array

	arr2 := []int{1,2,3,4,5,6}
	reverse(arr2)
	fmt.Println(arr2)

	// recursive way to reverse an array

	arr3 := []int{1,2,3,4,5,6}
	arr4 := recursiveReverse(arr3,0,len(arr3)-1)
	fmt.Println("Recursive reverse ",arr4)

	// frequency of each element in an array

	arr5 := []int{1,23,12,1,223,1,23}
	frequency(arr5)

	// frequency using map

	arr6 := []int{4,23,1,1,34,5,4}
	hashFrequency(arr6)

	// Rearrange increasing-decreasing order
	arr7 := []int{4,2,8,6,15,5,9,20}
	increasingDecreasing(arr7)
	fmt.Println(arr7)

	// Rotate array by k elements
	arr8 := []int{1,2,3,4,5,6,7}
	k :=3
	rotate(arr8,k)
	fmt.Println(arr8)
}

// second smallest element in an array
func second_smallest(arr []int, smallest int, sSmallest int) (int,int)  {

	smallest = arr[0]

	for i:=1; i<len(arr); i++ {

		if arr[i] < smallest {
			sSmallest = smallest
			smallest = arr[i]
		} else if arr[i] < sSmallest && arr[i] != smallest {
			sSmallest = arr[i]
		}
	}

	return smallest,sSmallest
}

// Reverse an array

func reverse(arr2 []int) {

	for i,j:=0,len(arr2)-1; i < j; i,j = i+1,j-1 {
		arr2[i],arr2[j] = arr2[j],arr2[i]
	}
}

// Recursive reverse


func recursiveReverse(arr3 []int,s int,l int) []int {

	if s > l {

		return arr3

	}

	arr3[s],arr3[l] = arr3[l],arr3[s]
	return recursiveReverse(arr3,s+1,l-1)

}

// frequency

func frequency(arr []int) {

	brr := make([]int,len(arr))
	var count int
	for i:=0; i<len(arr); i++ {
		count =1
		for j:=i+1; j<len(arr); j++ {

			if arr[i] == arr[j] && brr[j]!=-1 {
				brr[j]  = -1
				count++
			}
			
		}

		if brr[i]!= -1 {
			brr[i] = count
		}
	}

	for i:=0; i<len(arr); i++ {
		if brr[i] != -1 {
			fmt.Println(arr[i]," ",brr[i])
		}
		
	}
}

// Hash frequency

func hashFrequency(arr6 []int) {

	brr := make(map[int]int)

	for i:=0; i< len(arr6); i++ {

		_,ok := brr[arr6[i]];

		if !ok {
			brr[arr6[i]] = 1
		} else {
			brr[arr6[i]] = brr[arr6[i]] + 1
		}
	}

	fmt.Println(brr)
}


// Increasing Decreasing

func increasingDecreasing(arr []int) {

	mid := (0 + len(arr)-1)/2
	fmt.Println(mid)

	var k int
	var pos int

	for i:=0; i<len(arr); i++ {
		k = arr[i]
		pos = i

		for j := i+1; j<len(arr); j++ {

				if i <=mid {
					if arr[j] < k {
						k = arr[j]
						pos = j
					} 
				} else {
					if arr[j] > k {
						k = arr[j]
						pos = j
					}
				}
		}

		arr[i],arr[pos] = arr[pos],arr[i]
	}

	fmt.Println(arr)
}

// Rotate k items

func rotate(arr8 []int, k int) {

	for k > 0 {

		temp := arr8[0]
		for i:=0; i < len(arr8)-1; i++ {
			arr8[i] = arr8[i+1]

		}
		arr8[len(arr8)-1] = temp
		k--
	}
}