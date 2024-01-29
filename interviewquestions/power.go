package main

import "fmt"

func power(a, b int) int {
	if b == 0 {
		return 1
	}
	if b == 1 {
		return a
	}
	if b%2 == 1 {
		return power(a, b/2) * power(a, b/2) * a
	}
	return power(a, b/2) * power(a, b/2)
}

func main() {
	if power(2, 10) != 1024 {
		fmt.Println("Test 1 failed")
	}
	if power(2, 3) != 8 {
		fmt.Println("Test 2 failed")
	}
	if power(2, 8) != 256 {
		fmt.Println("Test 3 failed")
	}
	fmt.Println("Done")
}
