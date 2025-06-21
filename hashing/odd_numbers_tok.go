package main

import "fmt"

func odd_number_equals_to_k(nums []int, k int) {
	var m = make(map[int]int)
	m[0] = 1
	var count int
	var ans int
	for _, i := range nums {
		count += i % 2
		ans += m[count-k]
		m[count]++
	}
	fmt.Println(ans)
}
func main() {
	var nums = []int{1, 1, 2, 1, 1}
	var k = 3
	odd_number_equals_to_k(nums, k)
}
