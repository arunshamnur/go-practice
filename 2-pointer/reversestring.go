package main

import (
	"fmt"
)

func main() {
	var word = "abcdefd"
	var ch = []byte("d")[0]
	var runeword = []byte(word)

	for j := 1; j < len(runeword); j++ {
		if runeword[j] == ch {
			p := 0
			q := j
			for p < q {
				runeword[p], runeword[q] = runeword[q], runeword[p]
				p++
				q--
			}
			break
		}
	}
	fmt.Println(string(runeword))
}
