package main

import (
	"fmt"
)

func all_char_equal_occur(arr string) {
	var m = make(map[string]int)
	for i := 0; i < len(arr); i++ {
		m[string(arr[i])]++
	}
	for _, v := range m {
		if v != 2 {
			fmt.Println(false)
			return
		}
	}
	fmt.Println(true)
}

func main() {
	var s = "abacbc"
	all_char_equal_occur(s)
}
