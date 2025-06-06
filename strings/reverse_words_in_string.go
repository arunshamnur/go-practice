package main

import (
	"fmt"
	"strings"
)

func main() {
	var s = "Let's take LeetCode contest"
	s1 := strings.Split(s, " ")
	for i := range s1 {
		runes := []rune(s1[i]) // Convert to rune slice for Unicode safety
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i] // Swap characters
		}
		s1[i] = string(runes)
	}
	fmt.Println(strings.Join(s1, " "))

}
