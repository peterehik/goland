package main

import (
	"fmt"
	"strings"
)

func minLengthSubstring(s string, t string) int {

	//loop through s, for each element we check to see if that value is t and if it is, we remove it from t
	// by creating a new string that has all the elements in t except for the value we found and once new string is empty
	// we're done
	// Write your code here
	result := 0
	for i := 0; i < len(s); i++ {
		result += 1
		cur := s[i : i+1]
		indexOfCur := strings.Index(t, cur)
		if indexOfCur == -1 {
			continue
		}
		t = t[0:indexOfCur] + t[indexOfCur+1:] //remove cur from t
		if t == "" {
			break
		}
	}
	if t == "" {
		return result
	}
	return -1
}

func main() {
	// Call minLengthSubstring() with test cases here
	if minLengthSubstring("dcbefebce", "fd") != 5 {
		fmt.Println("Case 1 failed")
	}
	if minLengthSubstring("dcbefebce", "fde") != 5 {
		fmt.Println("Case 2 failed")
	}
	if minLengthSubstring("dcbefebce", "fdz") != -1 {
		fmt.Println("Case 3 failed")
	}
	fmt.Printf("Done")
}
