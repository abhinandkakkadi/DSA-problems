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

// quick sort

func QuickSort(arr []int,l int,h int) {

	if l < h {
		pivot := partition(arr,l,h) 
		QuickSort(arr,l,pivot-1)
		QuickSort(arr,pivot+1,h)
	}
}

func partition(arr []int, l int, h int) int {

	pivot := arr[h]
	i := l - 1

	for j:=l; j<h; j++ {
		
		if arr[j] < pivot {
			i++
			arr[i],arr[j] = arr[j],arr[i]
		}
	}

	arr[i+1],arr[h] = arr[h],arr[i+1]

	return i + 1
}
