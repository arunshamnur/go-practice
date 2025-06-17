package main

import (
	"fmt"
	"strconv"
)

func countelements(a []int) {
	var m = make(map[string]int)
	for i := 0; i < len(a); i++ {
		if _, ok := m[strconv.Itoa(a[i]-1)]; ok {
			fmt.Println(a[i] - 1)
		}
		m[strconv.Itoa(a[i])] = 1
	}
	fmt.Println(m)
}

func main() {
	var i = []int{1, 1, 3, 3, 5, 5, 7, 7}
	countelements(i)
}
