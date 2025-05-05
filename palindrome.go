package main

import "fmt"

func main() {
	var num = 121
	var temp = num
	var rev int
	for num > 0 {
		rem := num % 10
		rev = (rev * 10) + rem
		num = num / 10
	}
	if temp == rev {
		fmt.Println("num is palindrome")
	} else {
		fmt.Println("num is not a palindrome")
	}

}
