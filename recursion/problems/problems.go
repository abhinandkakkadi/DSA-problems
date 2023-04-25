package problems

import "fmt"

// print name n times
func Print(n int) {

	if n == 0 {
		return
	}

	fmt.Println("abhinand")
	Print(n-1)
}

// print 1 to n
func PrintNumbers(i int,n int) {

	if i == n + 1 {
		return
	}

	fmt.Print(i," ")
	PrintNumbers(i+1,n)
}

// print n to 1
func PrintNumbersReverse(n int) {
	
	if n == 0 {
		return
	}

	fmt.Print(n, " ")
	PrintNumbersReverse(n-1)
}

// sum of first n natural numbers
func Sum(n int) int {

	if n == 0 {
		return 0
	}

	return n + Sum(n-1)
}

// factorial of a number
func Fact(n int) int {

	if n == 1 || n == 0 {
		return n
	}

	return n * Fact(n-1)
}

// reverse an array
func ReverseArray(l int,h int,arr []int) {

	if l < h {
		arr[l],arr[h] = arr[h],arr[l]
		ReverseArray(l+1,h-1,arr)
	}

}

// palindrome using recursion
func Palindrome(l int,h int,s string) bool {

	if l > h {
		return true
	}

	if s[l] != s[h] {
		return false
	}

	return Palindrome(l+1,h-1,s)
}
