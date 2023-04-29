package main

import "fmt"



func main() {

	arr := []int{1,1,2,3,4,5,6,7,8}
	m := make(map[int]int)

	for i := range arr {
		m[arr[i]]++
	}

	for key,val := range m {
		fmt.Println(key," -> ",val)
	}
}