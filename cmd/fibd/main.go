package main

import (
	"fmt"
	"log"
	"os"
)

type Rabbits struct {
	mature uint64
	babies uint64
}

type Generation []Rabbits

func (g Generation) Before(i int) Rabbits {
	return g[len(g)-i]
}

func main() {
	content, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatalf("failed to open input file: %v", err)
	}

	var n, k int
	_, err = fmt.Sscanf(string(content), "%d %d", &n, &k)
	if err != nil {
		log.Fatalf("failed to extract final generation number and lifespan months: %v", err)
	}

	if n < 3 {
		fmt.Println("1")
		return
	}

	g := Generation{Rabbits{0, 1}, Rabbits{1, 0}}

	for range n - 2 {
		lastGen := g.Before(1)
		matureNow := lastGen.babies + lastGen.mature
		babiesNow := lastGen.mature

		if k <= len(g) {
			kThLast := g.Before(k)
			matureNow -= kThLast.babies
		}

		g = append(g, Rabbits{matureNow, babiesNow})
	}

	last := g[len(g)-1]
	fmt.Println(last.babies + last.mature)
}
