package main

import (
	"dsa/recursion/problems"
	"fmt"
)

func main() {

	// recursion to printf our name n times
	problems.Print(5)

	// print 1 to n using recursion
	problems.PrintNumbers(1,5)
	fmt.Println()

	// print n to 1 using recursion
	problems.PrintNumbersReverse(5)
	fmt.Println()

	// sum of first n natural numbers
	fmt.Println(problems.Sum(5))

	// factorial of a number
	fmt.Println(problems.Fact(5))

	// Reverse an array using recursion
	arr := []int{1,2,3,4,5,6}
	problems.ReverseArray(0,len(arr)-1,arr)
	fmt.Println(arr)


	// check whether a string is palindrome or not
	s := "malayalam"
	fmt.Println(problems.Palindrome(0,len(s)-1,s))
}