package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

const (
	MatchedLeft  = 0b01
	MatchedRight = 0b10
)

func main() {
	content, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatalf("failed to open input file: %v", err)
	}

	var s strings.Builder
	var candidates []string
	for chunk := range bytes.SplitSeq(content, []byte{'>'}) {
		labelSeen := false
		if len(chunk) == 0 {
			continue
		}
		for line := range bytes.Lines(chunk) {
			if !labelSeen {
				labelSeen = true
				continue
			}
			s.Write(bytes.TrimSpace(line))

		}

		candidates = append(candidates, s.String())
		s.Reset()
	}
	log.Println(len(candidates))

	adj := make(map[int]Chunk)
	pairs := make([]int, len(candidates))

	for i, a := range candidates {
		for j, b := range candidates {
			if i == j {
				continue
			}

			l := KMPOverlap(a, b)
			if l > 0 {
				adj[i] = Chunk{j, l}
				pairs[i] |= MatchedRight
				pairs[j] |= MatchedLeft
			}

		}
	}

	log.Println("kmp done", pairs)

	start := slices.Index(pairs, MatchedRight)
	s.Reset()
	s.WriteString(candidates[start])
	for chunk, ok := adj[start]; ok; chunk, ok = adj[chunk.next] {
		s.WriteString(candidates[chunk.next][chunk.size:])
	}
	fmt.Println(s.String())
}

type Chunk struct {
	next int
	size int
}

func KMPOverlap(a, b string) int {
	concat := b + "#" + a
	pi := make([]int, len(concat))
	n := len(concat)

	for i := 1; i < n; i++ {
		j := pi[i-1]
		for j > 0 && concat[i] != concat[j] {
			j = pi[j-1]
		}
		if concat[i] == concat[j] {
			j++
		}
		pi[i] = j
	}
	result := pi[n-1]
	if result < len(a)/2 || result < len(b)/2 {
		return 0
	}
	return result
}
