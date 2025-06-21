package main

import "fmt"

func bubblesort(a []int) {
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
}
func main() {
	var a = []int{8, 2, 6, 5, 7, 3}
	bubblesort(a)
	fmt.Println(a)
}
