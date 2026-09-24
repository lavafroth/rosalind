package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

func main() {
	content, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatalf("failed to open input file: %v", err)
	}

	Bases := []byte{'A', 'C', 'G', 'T'}
	var consensus [][]uint32 = nil
	for entry := range bytes.SplitSeq(content, []byte{'>'}) {
		if len(entry) == 0 {
			continue
		}
		_, seq, ok := bytes.Cut(entry, []byte{'\n'})
		if !ok {
			log.Fatal("unable to split fasta entry into label and sequences")
		}

		if consensus == nil {
			seqSize := 0
			for _, ch := range seq {
				if slices.Index(Bases, ch) != -1 {
					seqSize++
				}
			}

			for range Bases {
				consensus = append(consensus, make([]uint32, seqSize))
			}
		}

		y := 0
		for _, ch := range seq {
			x := slices.Index(Bases, ch)
			if x == -1 {
				continue
			}
			consensus[x][y] += 1
			y++
		}
	}

	var profile strings.Builder
	for j := range len(consensus[0]) {
		var maxVal uint32 = 0
		var maxBase byte = 'A'
		for i, base := range Bases {
			if consensus[i][j] > maxVal {
				maxVal = consensus[i][j]
				maxBase = base
			}
		}

		profile.WriteByte(maxBase)
	}
	fmt.Println(profile.String())
	for i, base := range Bases {
		repr := fmt.Sprint(consensus[i])
		fmt.Printf("%c: %v\n", base, repr[1:len(repr)-1])
	}

}
