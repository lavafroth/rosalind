package main

import (
	"fmt"
	"github.com/lavafroth/rosalind/revcomp"
	"log"
	"os"
)

func main() {
	content, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatalf("failed to open input file: %v", err)
	}

	r, err := revcomp.ReverseComplement(content)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(r))
}
