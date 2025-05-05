package main

import "fmt"

type stack []string

func (s *stack) push(key string) {
	*s = append(*s, key)
}

func (s *stack) isempty() bool {
	return len(*s) < 0
}

func (s *stack) pop() (string, bool) {
	if s.isempty() {
		return "", false
	} else {
		element := (*s)[len(*s)-1]
		*s = (*s)[:len(*s)-1]
		return element, true
	}
}

func main() {
	var s stack
	s.push("10")
	s.push("9")
	s.push("8")
	s.push("7")
	for len(s) > 0 {
		ele, status := s.pop()
		if status {
			fmt.Println(ele + " is removed")
		}
	}
}
