package main

import (
	"fmt"
	"github.com/gammazero/deque"
)

func main() {
	stack := deque.Deque[string]{}
	stack.PushFront("something")
	stack.PushFront("fuck off")

	fmt.Println(stack.PopFront())
	fmt.Println(stack)
}
