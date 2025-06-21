package main

import "fmt"

func mergesort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := mergesort(arr[:mid])
	right := mergesort(arr[mid:])
	return merge(left, right)
}

func merge(l []int, r []int) []int {
	var res []int
	var i, j int
	for i < len(l) && j < len(r) {
		if l[i] > r[j] {
			res = append(res, l[i])
			i++
		} else {
			res = append(res, r[j])
			j++
		}
	}
	for i < len(l) {
		res = append(res, l[i])
		i++
	}
	for j < len(r) {
		res = append(res, r[j])
		j++
	}

	return res
}

func main() {
	nums := []int{9, 5, 1, 4, 3}
	s := mergesort(nums)
	fmt.Println(s)
}
