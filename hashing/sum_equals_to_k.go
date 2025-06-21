package main

import "fmt"

func sum_equals_to_k(nums []int, k int) {
	var m = make(map[int]int)
	m[0] = 1
	var prefix_sum = 0
	var count int
	for i := 0; i < len(nums); i++ {
		prefix_sum += nums[i]
		if val, ok := m[prefix_sum-k]; ok {
			count += val
		}
		m[prefix_sum] += 1
	}
	fmt.Println(m)
	fmt.Println(count)
}
func main() {
	var nums = []int{1, -1, 0}
	var k = 0
	sum_equals_to_k(nums, k)
}
