package main

import "fmt"

func findMaxLength(nums []int) int {
	var prefixmap = make(map[int]int)
	var sum int
	prefixmap[0] = -1
	maxlen := 0
	for i, num := range nums {
		if num == 0 {
			sum += -1
		} else if num == 1 {
			sum += 1
		}
		if prev, ok := prefixmap[sum]; ok {
			curr := i - prev
			if curr > maxlen {
				maxlen = curr
			}
		}
		prefixmap[sum] = i
	}
	return maxlen
}

func main() {
	var arr = []int{0, 1, 1, 1, 1, 1, 0, 0, 0}
	fmt.Println(findMaxLength(arr))
}
