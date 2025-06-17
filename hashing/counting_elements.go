package main

import (
	"fmt"
	"strconv"
)

func countelements(arr []int) {
	var m = make(map[string]int)
	var count int
	for i := 0; i < len(arr); i++ {
		m[strconv.Itoa(arr[i])] = 1
	}

	for i := 0; i < len(arr); i++ {
		if _, ok := m[strconv.Itoa(arr[i]+1)]; ok {
			count++
		}
	}
	fmt.Println(count)
	fmt.Println(m)
}

func main() {
	var i = []int{1, 3, 2, 3, 5, 0}
	countelements(i)
}
