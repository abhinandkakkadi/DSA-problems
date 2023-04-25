package algorithm

// selection sort
func SelectionSort(arr []int, size int) {
	var small int
	var pos int
	for i := 0; i < size; i++  {
			small = arr[i]
			pos = i;

			for j := i + 1; j < size; j++ {
				
				if small > arr[j] {
					small = arr[j]
					pos = j
				}

			}

			arr[i],arr[pos] = arr[pos],arr[i]
	}

}

// buble sort
func BubbleSort(arr []int,size int) {
	var swap int
	for i :=0; i < size; i++ {
		swap = 0;
		for j := 0; j < size - 1 - i; j++ {
			if arr[j] > arr[j+1] {
				swap = 1
				arr[j],arr[j+1] = arr[j+1],arr[j]
			}
		}
		if swap == 0 {
			break
		}
	}
}

// insertion sort
func InsertionSort(arr []int, size int) {

	for i:=1; i < size; i++ {
		j := i - 1
		current := arr[i]
		for j >=0 && arr[j] > current {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = current
	}

}
