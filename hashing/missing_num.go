package main

import "fmt"

func check(nums []int) int {
	var m = make(map[int]int)
	length := len(nums)
	for i := 0; i < len(nums); i++ {
		if _, ok := m[nums[i]]; !ok {
			m[nums[i]]++
		}
	}
	for i := 0; i <= length; i++ {
		if _, ok := m[i]; !ok {
			return i
		}
	}
	return -1
}

func main() {
	var nums = []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
	fmt.Println(check(nums))
}
