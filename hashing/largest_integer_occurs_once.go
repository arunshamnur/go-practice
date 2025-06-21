package main

import "fmt"

func main() {
	var nums = []int{5, 7, 3, 9, 4, 9, 8, 3, 1}
	var m = make(map[int]int)
	var largest int
	var count int
	for _, v := range nums {
		m[v]++
	}
	for k, v := range m {
		if v == 1 {
			count++
			if k > largest {
				largest = k
			}
		}
	}
	if count == 0 {
		fmt.Println(-1)
	}

	fmt.Println(largest)
}
