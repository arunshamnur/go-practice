package main

func (l *LL) reverseLL() {
	//The idea is to reverse the links of all nodes using three pointers:
	//prev: pointer to keep track of the previous node
	//curr: pointer to keep track of the current node
	//next: pointer to keep track of the next node
	var prev *Node
	var cur = l.head
	for cur != nil {
		var next = cur.link
		cur.link = prev
		prev = cur
		cur = next
	}
	l.head = prev
}

func main() {
	ll := new(LL)
	ll.head = &Node{data: 1}
	ll.insertToLL(2)
	ll.insertToLL(3)
	ll.insertToLL(4)
	ll.insertToLL(5)
	ll.PrintLL()
	ll.reverseLL()
	ll.PrintLL()
}
