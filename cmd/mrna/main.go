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

	possibleCodons := map[byte]uint32{
		'T': 4,
		'K': 2,
		'S': 6,
		'Y': 2,
		'P': 4,
		'R': 6,
		'M': 1,
		'V': 4,
		'L': 6,
		'W': 1,
		'H': 2,
		'N': 2,
		'A': 4,
		'C': 2,
		'F': 2,
		'E': 2,
		'I': 3,
		'Q': 2,
		'G': 4,
		'D': 2,
	}

	modulus := uint32(1_000_000)
	totalCodons := uint32(3)
	for i, x := range content {
		if i == 0 {
			continue
		}
		p, ok := possibleCodons[x]
		if !ok {
			continue
		}
		totalCodons = totalCodons * p % modulus
	}
	fmt.Println(totalCodons)

}
