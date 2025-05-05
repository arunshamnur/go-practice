package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Node struct {
	data int
	link *Node
}
type linkedl struct {
	head *Node
	len  int
}

func (ll *linkedl) insertatBeg(data int) {
	var node = Node{data: data}
	if ll.head == nil {
		ll.head = &node
		ll.len += 1
		return
	}
	current := ll.head
	for current.link != nil {
		current = current.link
		if current.link == ll.head {
			break
		}
	}
	var temp = ll.head
	ll.head = &node
	ll.head.link = temp
	current.link = ll.head
	ll.len += 1
}
func (ll *linkedl) printLL() {
	var llcontent []string

	current := ll.head
	for current.link != nil {
		llcontent = append(llcontent, strconv.Itoa(current.data))
		current = current.link
		if current == ll.head {
			break
		}
	}
	llcontent = append(llcontent, strconv.Itoa(ll.head.data))
	fmt.Println(strings.Join(llcontent[:], "->"))
}

func (ll *linkedl) insertatEnd(data int) {
	if ll.head == nil {
		ll.insertatBeg(data)
		ll.len += 1
		return
	}
	current := ll.head
	for current.link != nil {
		if current.link == ll.head {
			break
		}
		current = current.link
	}
	current.link = &Node{data: data, link: ll.head}
}

func (ll *linkedl) reverse() {
	if ll.head == nil {
		return
	}
	var current = ll.head
	//var head = ll.head
	var prev *Node
	var node *Node
	for current != nil {
		node = current.link
		current.link = prev
		prev = current
		current = node
	}
	ll.head = prev
}

func main() {
	var ll = linkedl{}
	ll.insertatEnd(70)
	ll.insertatEnd(60)
	ll.insertatEnd(50)
	ll.insertatBeg(80)
	ll.insertatBeg(90)
	ll.insertatBeg(10)
	ll.printLL()
	ll.reverse()
	ll.printLL()
}
