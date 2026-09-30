package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	content, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatalf("failed to open input file: %v", err)
	}

	contents := string(content)
	count, rest, ok := strings.Cut(contents, "\n")
	if !ok {
		log.Fatal("failed to locate the number of nodes")
	}
	n, err := strconv.Atoi(count)
	if err != nil {
		log.Fatalf("failed to read the number of nodes: %v", err)
	}

	var group = make([]int, n+1)
	for i := range n + 1 {
		group[i] = i
	}

	for line := range strings.Lines(rest) {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		var a, b int
		fmt.Sscanf(line, "%d %d", &a, &b)

		bGrp := group[b]
		aGrp := group[a]

		for i, v := range group {
			if v == bGrp {
				group[i] = aGrp // left group chosen as greedy
			}
		}

	}

	var counter = make(map[int]bool)
	for _, g := range group[1:] {
		counter[g] = true
	}

	fmt.Println(len(counter) - 1)

}
