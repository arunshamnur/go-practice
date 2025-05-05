package main

import "fmt"

func partition(a []int, l int, r int) int {
	var pivot = a[r]
	var i = l
	for j := l; j < r; j++ {
		if a[j] < pivot {
			a[i], a[j] = a[j], a[i]
			i += 1
		}
	}
	a[i], a[r] = a[r], a[i]
	return i
}

func quicksort(a []int, l int, r int) {
	if l < r {
		var p = partition(a, l, r)
		quicksort(a, l, p-1)
		quicksort(a, p+1, r)
	}
}

func main() {
	var a = []int{7, 6, 5, 4, 3, 2}
	quicksort(a, 0, len(a)-1)
	fmt.Println(a)
}
