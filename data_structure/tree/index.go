package main

import "fmt"

type node struct {
	left  *node
	right *node
	data  int
}

func (n *node) insert(data int) {
	if n == nil {
		return
	} else if data <= n.data {
		if n.left == nil {
			n.left = &node{
				data:  data,
				left:  nil, //these aren't necessary, but they are here to make the code more readable
				right: nil, //these return the node pointer to the new node as its default value, nil
			}
		} else {
			n.left.insert(data)
		}
	} else {
		if n.right == nil {
			n.right = &node{
				data:  data,
				left:  nil,
				right: nil,
			}
		} else {
			n.right.insert(data)
		}
	}
}

// func (n *node) insert(data int) {
// 	if data <= n.data {
// 		if n.left == nil {
// 			n.left = &node{data: data}
// 		} else {
// 			n.left.insert(data)
// 		}
// 	} else {
// 		if n.right == nil {
// 			n.right = &node{data: data}
// 		} else {
// 			n.right.insert(data)
// 		}
// 	}
// }

type binaryTree struct {
	root *node
}

func (b *binaryTree) insert(data int) {
	if b.root == nil {
		b.root = &node{
			data:  data,
			left:  nil,
			right: nil,
		}
	} else {
		b.root.insert(data)
	}
}

func print(node *node, space int, ch string) {
	if node == nil {
		return
	}
	for i := 0; i < space; i++ {
		fmt.Printf(" ")
	}

	fmt.Printf("%s%d\n", ch, node.data)
	print(node.left, space+5, "L")
	print(node.right, space+5, "R")
}

func main() {
	tree := &binaryTree{}
	tree.insert(10)
	tree.insert(5)
	tree.insert(15)
	tree.insert(3)
	tree.insert(7)
	tree.insert(13)
	tree.insert(17)
	print(tree.root, 0, "")
}
