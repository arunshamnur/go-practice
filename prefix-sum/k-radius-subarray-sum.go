//You are given a 0-indexed array nums of n integers, and an integer k.
//
//The k-radius average for a subarray of nums centered at some index i with the radius k is the average of all elements in nums between the indices i - k and i + k (inclusive). If there are less than k elements before or after the index i, then the k-radius average is -1.
//
//Build and return an array avgs of length n where avgs[i] is the k-radius average for the subarray centered at index i.
//
//The average of x elements is the sum of the x elements divided by x, using integer division. The integer division truncates toward zero, which means losing its fractional part.
//
//For example, the average of four elements 2, 3, 1, and 5 is (2 + 3 + 1 + 5) / 4 = 11 / 4 = 2.75, which truncates to 2.
//
//
//Example 1:
//
//
//Input: nums = [7,4,3,9,1,8,5,2,6], k = 3
//Output: [-1,-1,-1,5,4,4,-1,-1,-1]
//Explanation:
//- avg[0], avg[1], and avg[2] are -1 because there are less than k elements before each index.
//- The sum of the subarray centered at index 3 with radius 3 is: 7 + 4 + 3 + 9 + 1 + 8 + 5 = 37.
//  Using integer division, avg[3] = 37 / 7 = 5.
//- For the subarray centered at index 4, avg[4] = (4 + 3 + 9 + 1 + 8 + 5 + 2) / 7 = 4.
//- For the subarray centered at index 5, avg[5] = (3 + 9 + 1 + 8 + 5 + 2 + 6) / 7 = 4.
//- avg[6], avg[7], and avg[8] are -1 because there are less than k elements after each index.

package main

import "fmt"

func main() {
	var nums = []int{40527, 53696, 10730, 66491, 62141, 83909, 78635, 18560}
	var k = 2
	var target []int

	var final []int
	var currentsum = 0
	for i := 0; i < len(nums); i++ {
		currentsum = currentsum + nums[i]
		target = append(target, currentsum)
	}

	for i := 0; i < len(nums); i++ {
		if i >= k && len(nums)-i > k {
			if i > k {
				final = append(final, (target[i+k]-target[i-k-1])/(k+k+1))
			} else {
				final = append(final, (target[i+k])/(k+k+1))

			}
		} else {
			final = append(final, -1)
		}
	}
	fmt.Println(final)

}
