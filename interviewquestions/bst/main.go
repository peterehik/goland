package main

import "fmt"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type Bst struct {
	Head *Node
}

func NewBst(head int) *Bst {
	return &Bst{
		Head: &Node{
			Value: head,
		},
	}
}

func (bst *Bst) Insert(value int) error {
	return bst.insert(bst.Head, &Node{Value: value})
}

func (bst *Bst) insert(head *Node, newNode *Node) error {
	if newNode.Value == head.Value {
		return fmt.Errorf("cannot have duplicate values in BST. %d already exists in tree", newNode.Value)
	}

	if newNode.Value < head.Value {
		if head.Left == nil {
			head.Left = newNode
			return nil
		}
		return bst.insert(head.Left, newNode)
	}
	// greater than, insert in right side
	if head.Right == nil {
		head.Right = newNode
		return nil
	}
	return bst.insert(head.Right, newNode)

}

func (bst *Bst) InOrderTraverse() []int {
	return bst.inOrderTraverse(bst.Head, make([]int, 0))
}

func (bst *Bst) inOrderTraverse(node *Node, vals []int) []int {
	if node == nil {
		return vals
	}
	vals = bst.inOrderTraverse(node.Left, vals)
	vals = append(vals, node.Value)
	vals = bst.inOrderTraverse(node.Right, vals)
	return vals
}

func (bst *Bst) Len() int {
	return bst.len(bst.Head)
}

func (bst *Bst) len(node *Node) int {
	if node == nil {
		return 0
	}
	return 1 + bst.len(node.Left) + bst.len(node.Right)
}

//func (bst *Bst) areTreesSimilar(node *Node, otherTreeElements []int, currentIndex int) bool {
//	if node == nil {
//		return true
//	}
//	leftSideTreeSimilar := bst.areTreesSimilar(node.Left, otherTreeElements, currentIndex)
//	if !leftSideTreeSimilar {
//		return false
//	}
//	if otherTreeElements[currentIndex] != node.Value {
//		return false
//	}
//	return bst.areTreesSimilar(node.Right, otherTreeElements, currentIndex+1)
//}
//
//func (bst *Bst) AreTreesSimilar(otherTreeElements []int) bool {
//	return bst.areTreesSimilar(bst.Head, otherTreeElements, 0)
//}

func main() {
	bst := NewBst(5)
	_ = bst.Insert(4)
	_ = bst.Insert(6)
	_ = bst.Insert(8)
	_ = bst.Insert(10)
	fmt.Printf("BST values: %+v, length: %d\n", bst.InOrderTraverse(), bst.Len())
}
