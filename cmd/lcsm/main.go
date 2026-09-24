package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	content, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatalf("failed to open input file: %v", err)
	}

	var longest = make(map[string]bool)
	longestEmpty := true

	var s strings.Builder
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

		if longestEmpty {
			longestEmpty = false
			longest = make(map[string]bool)
			longest[s.String()] = true
			s.Reset()
			continue
		}

		maxCandLen := 0
		var maxCands = make(map[string]bool)
		for sub := range longest {
			cands := LongestCommonSubstring(sub, s.String())
			for _, cand := range cands {
				candLen := len(cand)
				if candLen > maxCandLen {
					maxCandLen = candLen
					maxCands = make(map[string]bool)
					maxCands[cand] = true
				}
				if candLen == maxCandLen {
					maxCands[cand] = true
				}
			}
		}
		longest = maxCands
		s.Reset()
	}
	for anyLongest := range longest {
		fmt.Println(anyLongest)
		break
	}
}

func LongestCommonSubstring(a, b string) []string {

	r := len(a)
	n := len(b)

	// L dynamically stores the longest common substring
	// seen so far in an r cross n matrix.
	var L [][]int
	for range r {
		L = append(L, make([]int, n))
	}

	z := 0
	var commonStartAt []string

	for i := range r {
		for j := range n {
			if a[i] != b[j] {
				continue
			}

			if i == 0 || j == 0 {
				L[i][j] = 1
			} else {
				L[i][j] = L[i-1][j-1] + 1
			}

			if L[i][j] > z {
				z = L[i][j]
				commonStartAt = []string{a[i-z+1 : i+1]}
			}

			if L[i][j] == z {
				commonStartAt = append(commonStartAt, a[i-z+1:i+1])
			}
		}
	}
	return commonStartAt
}
