package main

import (
	"bytes"
	"fmt"
	"log"
	"math/big"
	"os"
)

func DisjointEdges(n int64) *big.Int {
	p := big.NewInt(1)
	for i := range n {
		p.Mul(p, big.NewInt(i+1))
	}
	return p
}

func main() {
	content, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatalf("failed to open input file: %v", err)
	}

	var aCount, cCount int64 = 0, 0
	for line := range bytes.Lines(content) {
		if len(line) == 0 || line[0] == '>' {
			continue
		}
		for _, c := range line {
			if c == 'A' {
				aCount += 1
			}
			if c == 'C' {
				cCount += 1
			}
		}
	}

	aEdges := DisjointEdges(aCount)
	fmt.Println(aEdges.Mul(aEdges, DisjointEdges(cCount)))
}
