//Example 2: 2270. Number of Ways to Split Array
//
//Given an integer array nums, find the number of ways to split the array
//into two parts so that the first section has a sum greater than or equal
//to the sum of the second section. The second section should have at
//least one number.

package main

import "fmt"

func main() {
	var a = []int{10, 8, -8, 7}

	var prefix_sum = []int{a[0]}
	for i := 1; i < len(a); i++ {
		prefix_sum = append(prefix_sum, a[i]+prefix_sum[len(prefix_sum)-1])
	}

	var count = 0

	for i := 0; i < len(a); i++ {
		left_sum := prefix_sum[i]
		right_sum := prefix_sum[len(a)-1] - prefix_sum[i]
		if left_sum > right_sum {
			count++
		}
	}
	fmt.Println(count)
}
