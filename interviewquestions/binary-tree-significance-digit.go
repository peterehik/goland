/*
*

The candidate is given a binary tree, with integer values at each node. The candidate is expected to return the sum of the numbers represented by the values on each node in each root-to-leaf path. The level of the tree is the significance digit on the number. Please see the example below.

root=> 2

	    / \
	   3   4
	 /  \
	1    5

e.g result here = 231 + 235 + 24 = 490

= 2*10^curdepth + 3*10^curdepth-1 + 1*10^1
*/
package main

import (
	"fmt"
	"github.com/gammazero/deque"
	"math"
)

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func sumRootToLeaf(head *Node) int {
	curPathStack := deque.New[int]()
	rootToLeafPaths := make(map[int]bool)
	build(head, curPathStack, rootToLeafPaths)
	fmt.Println(rootToLeafPaths)
	result := 0
	for sum, _ := range rootToLeafPaths {
		result += sum
	}
	return result
}

func saveVisitedStack(visitedStack *deque.Deque[int], rootToLeafPaths map[int]bool) {
	sum := 0
	var values []int
	repeat := visitedStack.Len()
	for i := 0; i < repeat; i++ {
		curVal := visitedStack.PopFront()
		sum += curVal * int(math.Pow(10, float64(i)))
		values = append(values, curVal)
	}
	rootToLeafPaths[sum] = false
	for i := repeat - 1; i >= 0; i-- {
		visitedStack.PushFront(values[i])
	}
}

func build(node *Node, visitedStack *deque.Deque[int], rootToLeafPaths map[int]bool) {
	if node == nil {
		return
	}
	visitedStack.PushFront(node.Val)
	build(node.Left, visitedStack, rootToLeafPaths)
	build(node.Right, visitedStack, rootToLeafPaths)
	if node.Left == nil && node.Right == nil {
		saveVisitedStack(visitedStack, rootToLeafPaths)
	}
	visitedStack.PopFront()
}

func main() {
	treeHead := &Node{
		Val: 5,
		Left: &Node{
			Val: 2,
			Left: &Node{
				Val: 3,
			},
			Right: &Node{
				Val: 6,
			},
		},
		Right: &Node{
			Val: 4,
			Left: &Node{
				Val: 1,
			},
		},
	}

	fmt.Printf("Sum of root to leaf: %d\n", sumRootToLeaf(treeHead))
}
