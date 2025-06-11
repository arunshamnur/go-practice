//A pangram is a sentence where every letter of the English alphabet appears at least once.
//
//Given a string sentence containing only lowercase English letters, return true if sentence is a pangram, or false otherwise.

package main

import "fmt"

func checkpanagram(sentence string) bool {
	var m = make(map[string]int)
	var count int
	for i := 0; i < len(sentence); i++ {
		ch := sentence[i]
		if _, ok := m[string(ch)]; ok {
			continue
		}
		m[string(ch)] = 1
		count++
	}
	if count == 26 {
		return true
	} else {
		return false
	}
}

func main() {
	var nums = "thequickbrownfoxjumpsoverthelazydog"
	fmt.Println(checkpanagram(nums))
}
