package main

import "fmt"

func permutation(a []string) [][]string {
	var res [][]string
	if len(a) == 0 {
		return res
	}
	if len(a) == 1 {
		res = append(res, a)
		return res
	}
	var l [][]string
	for k, _ := range a {
		var f = a[k]
		var rem []string
		rem = append(rem, a[:k]...)
		rem = append(rem, a[k+1:]...)
		var q = permutation(rem)
		for _, i := range q {
			var temp []string
			temp = append(temp, f)
			temp = append(temp, i...)
			l = append(l, temp)
		}
	}
	return l
}
func main() {
	var input = []string{"a", "b", "c"}
	fmt.Println(permutation(input))

}
