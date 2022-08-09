package main

import "fmt"

type queue []string

func (q *queue) enqueue(item string) {
	*q = append(*q, item)
}

func (q *queue) dequeue() string {
	item := (*q)[0]
	*q = (*q)[1:]
	return item
}

func main() {
	q := make(queue, 0)
	q.enqueue("a")
	q.enqueue("b")
	q.enqueue("c")
	fmt.Println("List:")
	fmt.Println(q)

	fmt.Println("Dequeued:", q.dequeue())
	fmt.Println("List:")
	fmt.Println(q)
}
