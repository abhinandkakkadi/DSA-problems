package main

import "fmt"


func main() {

	str := "aannccv"
	
	for i :=0; i < len(str); i++ {
		count := 0

		for j := 0; j < len(str); j++ {
			if str[i] == str[j] && i != j {
				count = 1
				break
			}
		}

		if count == 0 {
			fmt.Printf("%c",str[i])
			fmt.Println()
		}
	}
		
}

// map implementation