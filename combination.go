package main

import "fmt"

func comb(a []string, n int) [][]string {
	var res [][]string
	if n == 0 {
		var temp []string
		res = append(res, temp)
		return res
	}
	for k, _ := range a {
		var f = a[k]
		var rem []string
		rem = append(rem, a[k+1:]...)
		var temp = comb(rem, n-1)
		for _, j := range temp {
			var temp1 []string
			temp1 = append(temp1, f)
			temp1 = append(temp1, j...)
			res = append(res, temp1)
		}

	}
	return res
}

func main() {
	var a = []string{"a", "b", "c"}
	fmt.Println(comb(a, 2))

}
