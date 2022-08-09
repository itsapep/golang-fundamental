package main

import "fmt"

type stack []string

func (s *stack) push(item string) { // push an item to the top of the stack
	*s = append(*s, item)
}

func (s *stack) pop() string { // pop the top of the stack
	// sz:=s.size()=len(*s)
	// func (s *stack) size() int {
	// 	return len(*s)
	// }
	item := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return item
}

func (s *stack) peek() string { // peek at the top of the stack
	return (*s)[len(*s)-1]
}

func main() {
	s := make(stack, 0)
	s.push("a")
	s.push("b")
	s.push("c")
	fmt.Println("List:")
	fmt.Println(s)

	fmt.Println("Popped:", s.pop())
	fmt.Println("List:")
	fmt.Println(s)

	fmt.Println("Peek:", s.peek())
	fmt.Println("List:")
	fmt.Println(s)
}
