package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"slices"
)

func main() {
	content, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal("failed to read input file")
	}

	_, seq, ok := bytes.Cut(content, []byte("\n"))
	if !ok {
		log.Fatal("no newline found after the fasta record identifier")
	}

	var cleanSeq, revComp []byte
	comp := map[byte]byte{'A': 'T', 'T': 'A', 'C': 'G', 'G': 'C'}

	for _, n := range seq {
		c, ok := comp[n]
		if !ok {
			continue
		}
		cleanSeq = append(cleanSeq, n)
		revComp = append(revComp, c)
	}

	slices.Reverse(revComp)

	N := len(cleanSeq)
	var mat [][]uint32

	for range N {
		mat = append(mat, make([]uint32, N))
	}

	for i, n := range cleanSeq {
		for j, c := range revComp {

			if n != c {
				continue
			}
			if i == 0 || j == 0 {
				mat[i][j] = 1
			} else {
				mat[i][j] = mat[i-1][j-1] + 1
			}
		}
	}

	// DEBUG: display current matrix mat
	// for _, i := range lcs {
	// 	fmt.Println(i)
	// }

	for i := range N {
		for j := range N {
			current := mat[i][j]

			if current == 0 {
				continue
			}

			toRight := N - 1 - j
			jumpY := toRight - i
			palindromeLength := jumpY + 1

			if palindromeLength < 4 || palindromeLength > 12 {
				continue
			}
			if mat[i+jumpY][j+jumpY] == current+uint32(jumpY) {
				fmt.Println(i+1, palindromeLength)
			}

		}
	}
}
