package main

import "fmt"

//            4
//         /    \
//      2         6
//    /   \     /   \
// 1       3    5    7
//

type binarynode struct {
	val       int
	leftnode  *binarynode
	rightnode *binarynode
}
type bt struct {
	root *binarynode
}

func insert(b *binarynode, element int) *binarynode {
	if b == nil {
		b = &binarynode{val: element}
		return b
	}
	if element < b.val {
		b.leftnode = insert(b.leftnode, element)
	} else {
		b.rightnode = insert(b.rightnode, element)
	}
	return b
}

func (b *binarynode) DFS() {
	if b == nil {
		return
	}
	b.leftnode.DFS()
	fmt.Printf("%d->", b.val)
	b.rightnode.DFS()
}

func (b *binarynode) inorder() {
	if b == nil {
		return
	}
	b.leftnode.inorder()
	fmt.Printf("%d->", b.val)
	b.rightnode.inorder()
}

func (b *binarynode) preorder() {
	if b == nil {
		return
	}
	fmt.Printf("%d->", b.val)
	b.leftnode.preorder()
	b.rightnode.preorder()
}

func (b *binarynode) postorder() {
	if b == nil {
		return
	}
	b.leftnode.postorder()
	b.rightnode.postorder()
	fmt.Printf("%d->", b.val)
}

func maxdepth(b *binarynode) int {
	if b == nil {
		return 0
	}
	left := maxdepth(b.leftnode)
	right := maxdepth(b.rightnode)
	return max(left, right) + 1
}

func main() {
	var b = bt{root: &binarynode{val: 0}}
	b.root = insert(b.root, 1)
	b.root = insert(b.root, 2)
	b.root = insert(b.root, 3)
	b.root = insert(b.root, 4)
	b.root = insert(b.root, 5)
	b.root = insert(b.root, 6)
	b.root.DFS()
	fmt.Println()
	b.root.inorder()
	fmt.Println()
	b.root.preorder()
	fmt.Println()
	b.root.postorder()
	fmt.Println()
	fmt.Println(maxdepth(b.root))

}
