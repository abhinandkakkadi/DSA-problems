package array


func LargestElement(arr []int) int {
	largest := arr[0]
	for i := 1 ; i < len(arr); i++ {
		if largest < arr[i] {
			largest = arr[i]
		}
	}
	return largest
}

func SecondLargest(arr []int) int {

	var largest int
	var secondLargest int

	if arr[0] > arr[1] {
		largest = arr[0]
		secondLargest = arr[1]
	} else {
		largest = arr[1]
		secondLargest = arr[0]
	}

	for i := 2; i < len(arr); i++ {
		
		if arr[i] > largest {

			secondLargest = largest
			largest = arr[i]

		} else if arr[i] > secondLargest {

			secondLargest = arr[i]
			
		}
	}

	return secondLargest

}

// Remove duplicates from sorted array

func RemoveDuplicatesFromSortedArray(arr []int) int {

		temp := arr[0]
		j := 0

		for i := 1 ; i < len(arr); i++ {

			if arr[i] != temp {
				j++
				arr[j] = arr[i]
				temp = arr[j]
			}
		}

		return j+1
}

// left rotate array

func LeftRotate(arr []int, n int) {
	
	if n > len(arr)-1 {
		n = n%len(arr)-1
	}

	for n > 0 {

		temp := arr[0]
		for i :=0; i < len(arr)-1 ; i++ {
			arr[i] = arr[i+1]
		}	
		n--
		arr[len(arr)-1] = temp

	}
}

// Move zeroes to end

func MoveZeroes(arr []int) {

	j := -1

	for i := 0; i < len(arr); i++ {

		if arr[i] != 0 {
			j++
			arr[j],arr[i] = arr[i],arr[j]
		}
		
	}
}