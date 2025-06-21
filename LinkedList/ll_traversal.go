package main

import "fmt"

type Node struct {
	data int
	link *Node
}

type LL struct {
	head   *Node
	length int
}

func (n *LL) insert(data int) {
	var node = Node{data: data}
	if n.head == nil {
		n.head = &node
	} else {
		var cur = n.head
		for cur.link != nil {
			cur = cur.link
		}
		cur.link = &node
	}
}

func (n *LL) print() {
	if n.head == nil {
		return
	} else {
		var cur = n.head
		for cur != nil {
			fmt.Printf("%d->", cur.data)
			cur = cur.link
		}
	}
}

func (n *LL) reverse() {
	var prev *Node
	var curr = n.head
	for curr != nil {
		link := curr.link
		curr.link = prev
		prev = curr
		curr = link
	}
	n.head = prev
}

func (n *LL) addtostart(data int) {
	var node = Node{data: data}
	if n.head == nil {
		n.head = &node
	} else {
		var cur = n.head
		n.head = &node
		node.link = cur
	}
}

func (n *LL) addtoend(data int) {
	var node = Node{data: data}
	if n.head == nil {
		n.head = &node
	} else {
		cur := n.head
		for cur.link != nil {
			cur = cur.link
		}
		cur.link = &node
	}
}

func (n *LL) deltostart() {
	if n.head == nil {
		return
	} else {
		var cur = n.head
		n.head = cur.link
		cur.link = nil
	}

}

func removeduplicated(head *Node) *Node {
	if head == nil {
		return nil
	}
	cur := head
	for cur.link != nil {
		if cur.data == cur.link.data {
			cur.link = cur.link.link
		} else {
			cur = cur.link
		}
	}

	return head
}

func (n *LL) deltogivenindex(index int) {
	if n.head == nil {
		return
	} else {
		var count int
		var cur = n.head
		for cur.link != nil {
			count++
			if count+1 == index {
				cur.link = cur.link.link
			}
			cur = cur.link
		}
	}
}

func (n *LL) deltoend() {
	if n.head == nil {
		return
	} else {
		var cur = n.head
		for cur.link.link != nil {
			cur = cur.link
		}
		cur.link = nil
	}
}

func (n *LL) addtogivenindex(data int, index int) {
	var node = Node{data: data}
	if n.head == nil {
		n.head = &node
	} else {
		var count int
		var cur = n.head
		for cur.link != nil {
			count++
			if count == index-1 {
				link := cur.link
				cur.link = &node
				node.link = link
				break
			}
			cur = cur.link
		}
	}
}
func main() {
	l := &LL{}
	l.addtoend(2)
	l.addtoend(4)
	l.addtostart(1)
	l.addtogivenindex(3, 3)
	l.addtoend(5)
	l.addtoend(6)
	l.print()
	fmt.Println()
	//l.deltostart()
	//l.deltoend()
	//l.print()
	l.deltogivenindex(2)
	fmt.Println()
	l.print()
	//l.reverse()
	//fmt.Println()
	//l.print()
}
