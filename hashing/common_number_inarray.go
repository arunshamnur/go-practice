package main

import (
	"fmt"
)

func intersection_of_multiple_array(arr [][]int) {
	var m = make(map[int]int)
	//var a = make([]int, 0)
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			m[arr[i][j]]++
		}
	}
	for k, v := range m {
		if v == len(arr) {
			fmt.Println(k)
		}
	}
}

func main() {
	var a = [][]int{{3, 1, 2, 4, 5}, {1, 2, 3, 4}, {3, 4, 5, 6}}
	intersection_of_multiple_array(a)
}
