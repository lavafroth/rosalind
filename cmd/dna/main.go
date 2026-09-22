package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	content, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatalf("failed to open input file: %v", err)
	}

	count := make(map[byte]uint32)
	for _, c := range content {
		count[c] += 1
	}

	for i, c := range []byte{'A', 'C', 'G', 'T'} {
		sep := " "
		if i == 3 {
			sep = "\n"
		}
		fmt.Print(count[c], sep)
	}

}
