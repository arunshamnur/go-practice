package main

import "fmt"

type binaryNode struct {
	data      int
	leftNode  *binaryNode
	rightNode *binaryNode
}
type binaryTree struct {
	root *binaryNode
	len  int
}

func insert(bnode *binaryNode, data int) *binaryNode {
	var node = binaryNode{data: data}
	if bnode == nil {
		return &node
	} else {
		if bnode.data < data {
			bnode.rightNode = insert(bnode.rightNode, data)
		} else {
			bnode.leftNode = insert(bnode.leftNode, data)
		}
	}
	return bnode
}

func inorder(node *binaryNode) {
	if node != nil {
		inorder(node.leftNode)
		fmt.Printf("%d->", node.data)
		inorder(node.rightNode)
	}

}
func postorder(node *binaryNode) {
	if node != nil {
		inorder(node.leftNode)
		inorder(node.rightNode)
		fmt.Printf("%d->", node.data)
	}

}
func preorder(node *binaryNode) {
	if node != nil {
		fmt.Printf("%d->", node.data)
		inorder(node.leftNode)
		inorder(node.rightNode)
	}

}

func main() {
	var bt = binaryTree{}
	bt.root = insert(bt.root, 50)
	bt.root = insert(bt.root, 30)
	bt.root = insert(bt.root, 10)
	bt.root = insert(bt.root, 35)
	bt.root = insert(bt.root, 60)
	bt.root = insert(bt.root, 55)
	bt.root = insert(bt.root, 70)
	inorder(bt.root)
	fmt.Println()
	postorder(bt.root)
	fmt.Println()
	preorder(bt.root)

}
