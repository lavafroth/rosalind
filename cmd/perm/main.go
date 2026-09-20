package main

import (
	"fmt"
	"log"
	"os"
)

func display(a []uint32) {
	s := fmt.Sprint(a)
	fmt.Println(s[1 : len(s)-1])
}

func main() {
	content, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal("failed to read input file")
	}

	// N <= 7
	N := uint32(content[0] - '0')

	count := uint32(1)
	for i := range N {
		count *= i + 1
	}

	var a, p []uint32

	for x := range N {
		a = append(a, x+1)
		p = append(p, 0)
	}
	fmt.Println(count)
	display(a)
	for i := uint32(1); i < N; {
		if p[i] < i {

			j := i & 1 * p[i]
			a[j], a[i] = a[i], a[j]
			p[i] += 1
			i = 1
			display(a)
			continue
		}

		p[i] = 0
		i += 1
	}
}

// notes
// synteny block: A DNA block condensed into a unit for genomic comparison.
// genome rearrangements: rearrangement of entire intervals of DNA.
//
// We label each synteny block with a positive integer.
