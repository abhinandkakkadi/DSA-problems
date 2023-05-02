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

// merge sort

func MergeSort(arr []int,l int, h int) {

	if l < h {
		mid := (l+h)/2
		MergeSort(arr,l,mid)
		MergeSort(arr,mid+1,h)
		Merge(arr,l,mid,h)
	}
}

func Merge(arr []int, l int, mid int, h int) {

	i := l
	j := mid + 1
	var b []int
	for i <= mid && j <= h {
		if arr[i] <= arr[j] {
			b = append(b, arr[i])
			i++
		} else {
			b = append(b, arr[j])
			j++
		}
		
	}

	for ; i<=mid; i++ {
		b =append(b, arr[i])
	}

	for ; j <=h; j++ {
		b = append(b, arr[j])
	}

	for i:=l; i <=h; i++ {
		arr[i] = b[i-l]
	}
}