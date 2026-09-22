package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

func main() {
	content, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatalf("failed to open input file: %v", err)
	}

	slices.Reverse(content)

	var s strings.Builder
	for _, c := range content {
		switch c {
		case 'A':
			s.WriteByte('T')
		case 'T':
			s.WriteByte('A')
		case 'C':
			s.WriteByte('G')
		case 'G':
			s.WriteByte('C')
		}
	}

	fmt.Println(s.String())
}
