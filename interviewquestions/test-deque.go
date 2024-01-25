package main

import (
	"fmt"
	"github.com/gammazero/deque"
)

func main() {
	myQueue := deque.Deque[string]{}
	myQueue.PushFront("something")
	myQueue.PushFront("fuck off")
	
	fmt.Println(myQueue.PopFront())
	fmt.Println(myQueue)
}
