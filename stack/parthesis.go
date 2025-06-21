package main

import "fmt"

type stack []string

func (s *stack) push(element string) {
	*s = append(*s, element)
}

func (s *stack) isempty() bool {
	return len(*s) == 0
}

func (s *stack) pop() string {
	if s.isempty() {
		return ""
	} else {
		element := (*s)[len(*s)-1]
		*s = (*s)[:len(*s)-1]
		return element
	}
}

func checkvalid(s stack, inputstring string) bool {
	var m = make(map[string]string)
	m["("] = ")"
	m["["] = "]"
	m["{"] = "}"
	for _, element := range inputstring {
		if string(element) == "(" || string(element) == "{" || string(element) == "[" {
			s.push(string(element))
		} else {
			if m[s.pop()] != string(element) {
				return false
			}
		}
	}
	if s.isempty() {
		return true
	} else {
		return false
	}
}

func main() {
	var s stack
	var inputstring = "({})"

	fmt.Println(checkvalid(s, inputstring))

	// var s stack
	// s.push("A")
	// s.push("B")
	// s.push("C")
	// s.push("D")
	// fmt.Println(s.pop())
	// fmt.Println(s.pop())
	// fmt.Println(s.pop())
	// fmt.Println(s.pop())

}
