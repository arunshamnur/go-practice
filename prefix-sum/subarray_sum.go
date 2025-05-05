//Example 1: Given an integer array nums, an array queries where queries[i] = [x, y] and
//an integer limit, return a boolean array that represents the answer to each query.
//A query is true if the sum of the subarray from x to y is less than limit, or false otherwise.
//
//For example, given nums = [1, 6, 3, 2, 7, 2], queries = [[0, 3], [2, 5], [2, 4]], and limit = 13,
//the answer is [true, false, true]. For each query, the subarray sums are [12, 14, 12].
//

package main

import "fmt"

func main() {
	var a = []int{1, 6, 3, 2, 7, 2}

	var b = [][]int{{0, 3}, {2, 5}, {2, 4}}
	limit := 13
	var returnArr []bool

	var prefix_sum = []int{a[0]}
	for i := 1; i < len(a); i++ {
		prefix_sum = append(prefix_sum, a[i]+prefix_sum[len(prefix_sum)-1])
	}

	fmt.Println(prefix_sum)
	for i := 0; i < len(b); i++ {
		if b[i][0] > 0 {
			if (prefix_sum[b[i][1]] - prefix_sum[b[i][0]-1]) < limit {
				returnArr = append(returnArr, true)
			} else {
				returnArr = append(returnArr, false)
			}
		} else {
			if (prefix_sum[b[i][1]]) < limit {
				returnArr = append(returnArr, true)
			} else {
				returnArr = append(returnArr, false)
			}
		}
	}
	fmt.Println(returnArr)
}
