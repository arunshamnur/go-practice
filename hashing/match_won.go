package main

import (
	"fmt"
	"sort"
)

func main() {
	var arr = [][]int{{1, 3}, {2, 3}, {3, 6}, {5, 6}, {5, 7}, {4, 5}, {4, 8}, {4, 9}, {10, 4}, {10, 9}}
	var m = make(map[int]int)
	var ans = make([][]int, 2)
	for _, val := range arr {
		if _, ok := m[val[0]]; !ok {
			m[val[0]] = 0
		}
		m[val[1]]++
	}
	for k, val := range m {
		if val == 0 {
			ans[0] = append(ans[0], k)
		}
		if val == 1 {
			ans[1] = append(ans[1], k)
		}
	}
	sort.Ints(ans[0])
	sort.Ints(ans[1])
	fmt.Println(ans)
}
