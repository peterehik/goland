// Given those two arrays:
// A = [4, 2, 6]
// B = [5, 1, 3, 0, 0, 0]

// I would like to have a function that sorts and merges those two arrays into one array. I would like the solution to optimize for memory allocation.
// Assumptions:
// 1. Size of array B is always 2*size of array A.
// 2. Elements of the B array at the index that is greater than len(A) (which is 3 in this example) are place holders
// In this example, the result should be [1,2,3,4,5,6];

package main

import (
	"fmt"
	"sort"
)

func main() {
	TestMergeSort()
}

func MergeAndSort(a, b []int) []int {
	lenA := len(a)
	j := 0
	for i := lenA; i < len(b); i++ {
		b[i] = a[j]
		j++
	}
	sort.Slice(b, func(i, j int) bool {
		return b[i] <= b[j]
	})
	return b
}

func TestMergeSort() {
	testcases := []struct {
		name     string
		first    []int
		second   []int
		expected []int
	}{
		{
			name:     "test 1",
			first:    []int{4, 2, 6},
			second:   []int{5, 1, 3, 0, 0, 0},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:     "test 2",
			first:    []int{1, 2, 3},
			second:   []int{7, 6, 5, 0, 0, 0},
			expected: []int{1, 2, 3, 5, 6, 7},
		},
		// both of them are empty
		// they contain the same elements i..e [1,1, 1] [1,1,1,0,0,0]
		// 1 million records
	}
	for _, tc := range testcases {

		MergeAndSort(tc.first, tc.second)
		if len(tc.expected) != len(tc.second) {
			panic(fmt.Errorf("lengths are not equal: expected %d got %d", len(tc.expected), len(tc.second)))
		}

		for i, element := range tc.expected {
			if element != tc.second[i] {
				panic(fmt.Errorf("elements are not equal: expected %d got %d", element, tc.second[i]))
			}
		}
		fmt.Printf("Success %s\n", tc.name)
	}
}
