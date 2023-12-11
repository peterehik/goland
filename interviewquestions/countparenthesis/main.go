package main

import (
	"bufio"
	"fmt"
	"github.com/golang-collections/collections/stack"
	"os"
	"strings"
)

// CountParenthesis returns the number of opened parenthesis
// ((( returns 3
// (()) returns 0
// (((()) returns 2
// This question was from my facebook mock interview
// ((() (( )((((
func CountParenthesis(input string) (int, error) {
	s := stack.New()
	for _, char := range input {
		if char == '(' {
			s.Push(char)
		} else if char == ')' {
			lastChar := s.Peek()
			if lastChar == '(' {
				s.Pop()
			} else {
				s.Push(char)
			}
		} else {
			return 0, fmt.Errorf("found %c. input must only contain ( and ) chars", char)
		}
	}

	return s.Len(), nil
}

func main() {
	sampleInputs := []string{
		"(((((())))", "(((()", ")(()))",
	}
	for _, input := range sampleInputs {
		count, err := CountParenthesis(input)
		if err != nil {
			panic(err)
		}
		fmt.Printf("input: %s, result: %d\n", input, count)
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("Keep playing (Y|N): ")
		text, _ := reader.ReadString('\n')
		text = strings.Trim(text, "\n ")
		if text != "Y" {
			fmt.Printf("We're done here, thanks")
			break
		}
		fmt.Printf("Sample input: ")
		text, _ = reader.ReadString('\n')
		text = strings.Trim(text, "\n ")
		count, err := CountParenthesis(text)
		if err != nil {
			panic(err)
		}
		fmt.Printf("input: %s, result: %d\n", text, count)
	}
}
