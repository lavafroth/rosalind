package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatalf("failed to open input file: %v", err)
	}

	var count int
	var perm []int
	fmt.Fscanf(file, "%d", &count)
	for i := range count {
		var word string
		_, err := fmt.Fscan(file, &word)
		if err != nil {
			log.Fatalf("failed to read the %d-th element of the permutation: %v", i, err)
		}
		v, err := strconv.Atoi(word)
		if err != nil {
			log.Fatalf("failed to parse the %d-th element of the permutation: %v", i, err)
		}
		perm = append(perm, v)
	}

	longestSequence(true, perm)
	longestSequence(false, perm)
}

func longestSequence(ascending bool, seq []int) {
	n := len(seq)
	if n == 0 {
		return
	}

	var tops, parent, out []int
	for range n {
		parent = append(parent, -1)
	}

	for i, v := range seq {
		if i == 0 {
			tops = append(tops, 0) // first element
			continue
		}

		lo, hi := 0, len(tops)

		for hi > lo {
			mid := (lo + hi) / 2
			if ascending == (seq[tops[mid]] < v) {
				lo = mid + 1
			} else {
				hi = mid
			}
		}

		if lo < len(tops) {
			tops[lo] = i
		} else {
			tops = append(tops, i)
		}

		if lo > 0 {
			parent[i] = tops[lo-1]
		}
	}

	for i := tops[len(tops)-1]; i != -1; i = parent[i] {
		out = append(out, seq[i])
	}

	slices.Reverse(out)
	Display(out)

}

func Display(a []int) {
	for i, v := range a {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(v)
	}
	fmt.Println()
}
