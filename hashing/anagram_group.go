package main

import (
	"fmt"
	"sort"
	"strings"
)

func groupAnagrams(s []string) {
	var m = map[string][]string{}
	var res = make([][]string, 0)
	for _, str := range s {
		split := strings.Split(str, "")
		sort.Strings(split)
		res := strings.Join(split, "")
		m[res] = append(m[res], str)
	}
	for _, str := range m {
		res = append(res, str)
	}
	fmt.Println(res)
}

func main() {
	var strs = []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	groupAnagrams(strs)
}
