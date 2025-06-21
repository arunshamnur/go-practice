package main

import "fmt"

func main() {
	var arr = []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}
	var k = 2
	var i int
	var zero_count int
	var maxlength int
	var longestArr []int
	for j := 0; j < len(arr); j++ {
		if arr[j] == 0 {
			zero_count++
		}
		if zero_count > k {
			for zero_count > k && i < j {
				if arr[i] == 0 {
					zero_count--
				}
				i++
			}
		}
		if maxlength <= j-i+1 {
			maxlength = j - i + 1
			longestArr = arr[i : j+1]
		}
	}
	fmt.Println(maxlength)
	fmt.Println(longestArr)
}
