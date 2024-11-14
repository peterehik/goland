package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

var sampleParagraph = `"The quick brown fox jumps over the lazy dog. The quick fox is quick and smart. The dog, however, is lazy and sleepy."`

func countRepititions(paragraph string) int {
	lowercasepg := strings.ToLower(paragraph)
	wordRegex := regexp.MustCompile(`\w+`)
	matches := wordRegex.FindAllString(lowercasepg, -1)
	result := map[string]int{}
	for _, match := range matches {
		val, found := result[match]
		if !found {
			result[match] = 1
		} else {
			result[match] = val + 1
		}
	}

	numRepeatedWords := 0
	for _, count := range result {
		if count > 1 {
			numRepeatedWords++
		}
	}
	return numRepeatedWords
}

func main() {
	if countRepititions(sampleParagraph) == 7 {
		fmt.Println("Test passed!")
		return
	}
	log.Fatalln("Test failed!")
}
