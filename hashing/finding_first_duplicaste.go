package main

import "fmt"

func main() {
	var nums = "abcdeda"
	var m = make(map[string]int)
	var p, q int
	for i := 0; i < len(nums); i++ {
		ch := nums[i]
		if _, ok := m[string(ch)]; ok {
			p, q = m[string(ch)], i
			break
		}
		m[string(ch)] = i
	}
	fmt.Println(p, q, nums[p], nums[q])
	fmt.Println(m)
}
