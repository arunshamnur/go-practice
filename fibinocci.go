package main

import "fmt"

func main() {
	var num = 9
	var n1 = 0
	var n2 = 1
	fmt.Println(n1)
	fmt.Println(n2)
	n3 := n1 + n2
	for i := 3; i <= num; i++ {
		fmt.Println(n3)
		n1 = n2
		n2 = n3
		n3 = n1 + n2
	}

	fmt.Println("---------------------")
	for j := 0; j < num; j++ {
		fmt.Println(withrecur(j))
	}
}

func withrecur(num int) int {
	if num == 0 {
		return 0
	} else if num == 1 {
		return 1
	} else {
		return withrecur(num-1) + withrecur(num-2)
	}
}
