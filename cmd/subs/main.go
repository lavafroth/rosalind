package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {
	content, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatalf("failed to open input file: %v", err)
	}

	haystack, needle, ok := bytes.Cut(content, []byte{'\n'})
	if !ok {
		log.Fatal("failed to split input into haystack and needle")
	}

	haystack = bytes.TrimSpace(haystack)
	needle = bytes.TrimSpace(needle)
	windows := len(haystack) - len(needle) + 1

	var output []int

	for shift := range windows {
		start := shift
		stop := shift + len(needle)
		if bytes.Equal(haystack[start:stop], needle) {
			output = append(output, shift+1)
		}
	}

	repr := fmt.Sprint(output)
	fmt.Println(repr[1 : len(repr)-1])
}
