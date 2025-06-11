package main

import "fmt"

func main() {
	var nums = []int{5, 2, 7, 10, 3, 9}
	target := 8
	var p, q int
	var m = make(map[int]int)
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		rem := target - num
		if _, ok := m[rem]; ok {
			p, q = i, m[num]
			break
		}
		m[num] = i
	}
	fmt.Println(p, q, nums[p], nums[q])
	fmt.Println(m)
}
