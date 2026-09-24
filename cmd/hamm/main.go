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

	a, b, ok := bytes.Cut(content, []byte{'\n'})
	if !ok {
		log.Fatal("failed to separate two input sequences")
	}
	a = bytes.TrimSpace(a)
	b = bytes.TrimSpace(b)

	if len(a) != len(b) {
		log.Fatal("two input sequence are not of the same length")

	}

	hamm := 0

	for i := range len(a) {
		if a[i] != b[i] {
			hamm++
		}
	}

	fmt.Println(hamm)

}
