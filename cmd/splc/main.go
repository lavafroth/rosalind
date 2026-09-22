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
		log.Fatal("failed to read the input file")
	}

	var samples [][]byte

	for chunk := range bytes.SplitSeq(content, []byte(">")) {
		_, after, ok := bytes.Cut(chunk, []byte("\n"))
		if !ok {
			continue
		}

		sample := bytes.ReplaceAll(after, []byte("\n"), []byte{})
		samples = append(samples, sample)

	}

	dna := samples[0]
	introns := samples[1:]

	n := len(introns)
	matchCount := make([]int, n)

	ranges := []int{0}

	for i := range len(dna) {
		for j, intron := range introns {
			if intron[matchCount[j]] == dna[i] {
				matchCount[j] += 1
			} else if intron[0] == dna[i] {
				matchCount[j] = 1
			} else {
				matchCount[j] = 0
			}

			if matchCount[j] == len(intron) {
				stop := i + 1
				start := stop - matchCount[j]

				ranges = append(ranges, start)
				ranges = append(ranges, stop)
				for k := range n {
					matchCount[k] = 0
				}
			}
		}
	}

	ranges = append(ranges, len(dna))
	var exons strings.Builder
	for i := range len(ranges) / 2 {
		start := ranges[2*i]
		stop := ranges[2*i+1]
		exons.WriteString(string(dna[start:stop]))
	}

	codonProteins := map[string]byte{
		"TTT": 'F',
		"CTT": 'L',
		"ATT": 'I',
		"GTT": 'V',
		"TTC": 'F',
		"CTC": 'L',
		"ATC": 'I',
		"GTC": 'V',
		"TTA": 'L',
		"CTA": 'L',
		"ATA": 'I',
		"GTA": 'V',
		"TTG": 'L',
		"CTG": 'L',
		"ATG": 'M',
		"GTG": 'V',
		"TCT": 'S',
		"CCT": 'P',
		"ACT": 'T',
		"GCT": 'A',
		"TCC": 'S',
		"CCC": 'P',
		"ACC": 'T',
		"GCC": 'A',
		"TCA": 'S',
		"CCA": 'P',
		"ACA": 'T',
		"GCA": 'A',
		"TCG": 'S',
		"CCG": 'P',
		"ACG": 'T',
		"GCG": 'A',
		"TAT": 'Y',
		"CAT": 'H',
		"AAT": 'N',
		"GAT": 'D',
		"TAC": 'Y',
		"CAC": 'H',
		"AAC": 'N',
		"GAC": 'D',
		"CAA": 'Q',
		"AAA": 'K',
		"GAA": 'E',
		"CAG": 'Q',
		"AAG": 'K',
		"GAG": 'E',
		"TGT": 'C',
		"CGT": 'R',
		"AGT": 'S',
		"GGT": 'G',
		"TGC": 'C',
		"CGC": 'R',
		"AGC": 'S',
		"GGC": 'G',
		"CGA": 'R',
		"AGA": 'R',
		"GGA": 'G',
		"TGG": 'W',
		"CGG": 'R',
		"AGG": 'R',
		"GGG": 'G',
	}

	exonStr := exons.String()
	var protein strings.Builder

	for i := range len(exonStr) / 3 {
		codon := exonStr[3*i : 3*i+3]
		if codon == "TAA" || codon == "TAG" || codon == "TGA" {
			break
		}

		prot, ok := codonProteins[codon]
		if !ok {
			log.Fatalf("unknown codon: %s", codon)
		}
		protein.WriteByte(prot)
	}

	fmt.Println(protein.String())
}
