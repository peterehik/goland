// We don’t provide test cases in this language yet, but have outlined the signature for you. Please write your code below, and don’t forget to test edge cases!
/**
Given a sequence of n integers arr, determine the lexicographically smallest sequence which may be obtained from it after performing at most k element swaps, each involving a pair of consecutive elements in the sequence.
Note: A list x is lexicographically smaller than a different equal-length list y if and only if, for the earliest index at which the two lists differ, x's element at that index is smaller than y's element at that index.
Signature
int[] findMinArray(int[] arr, int k)
Input
n is in the range [1, 1000].
Each element of arr is in the range [1, 1,000,000].
k is in the range [1, 1000].
Output
Return an array of n integers output, the lexicographically smallest sequence achievable after at most k swaps.
Example 1
n = 3
k = 2
arr = [5, 3, 1]
output = [1, 5, 3]
We can swap the 2nd and 3rd elements, followed by the 1st and 2nd elements, to end up with the sequence [1, 5, 3]. This is the lexicographically smallest sequence achievable after at most 2 swaps.
Example 2
n = 5
k = 3
arr = [8, 9, 11, 2, 1]
output = [2, 8, 9, 11, 1]
We can swap [11, 2], followed by [9, 2], then [8, 2].
*/
package main

import (
	"fmt"
)

func findMinArray(arr []int, k int) []int {
	// Write your code here

	// go through array comparing element with next, if element is bigger than next swap them
	continueLooping := true
	for k > 0 && continueLooping {
		continueLooping = false

		for i := 0; i < len(arr)-1; i++ {
			if arr[i] > arr[i+1] {
				arr = swapElementsAt(arr, i, i+1)
				k--
				continueLooping = true
				break
			}
		}
	}
	return arr
}

func swapElementsAt(array []int, i, j int) []int {
	var tmp int
	tmp = array[i]
	array[i] = array[j]
	array[j] = tmp
	return array
}

func main() {

	fmt.Printf("%+v\n", findMinArray([]int{8, 9, 11, 2, 1}, 3))
	fmt.Printf("%+v\n", findMinArray([]int{5, 3, 1}, 2))
	fmt.Printf("%+v\n", findMinArray([]int{1, 2, 3, 4, 5}, 5))

	// Call findMinArray() with test cases here
}

/**

The candidate is given a binary tree, with integer values at each node. The candidate is expected to return the sum of the numbers represented by the values on each node in each root-to-leaf path. The level of the tree is the significance digit on the number. Please see the example below.

root=> 2
      / \
     3   4
   /  \
  1    5


Implement the power function that computes a to the power of b using basic arithmetic operations (+, -, *, /).

*/
