package main

import "fmt"

func longestsubstring(s string, k int) {
	var m = make(map[string]int)
	var p = 0
	var q int
	var ans []string
	for i := 0; i < len(s); i++ {
		m[string(s[i])]++
		q++
		for len(m) > k {
			if len(ans) > 0 {
				if len(ans[0]) < len(s[p:q-1]) {
					ans[0] = s[p : q-1]
				}
			} else {
				ans = append(ans, s[p:q-1])
			}
			m[string(s[p])]--
			if m[string(s[p])] == 0 {
				delete(m, string(s[p]))
			}
			p++
		}

		if len(s)-1 == i {
			if len(ans[0]) < len(s[p:q-1]) {
				ans[0] = s[p : q-1]
			}

		}
	}
	fmt.Println(ans[0])
	//fmt.Println(s[p : q+1])
}
func main() {
	var s = "eceba"
	longestsubstring(s, 2)
}
