package main

import (
	"fmt"
	"iter"
	"log"
	"os"
	"strings"
)

func Codons(b []byte) iter.Seq[[]byte] {
	return func(yield func([]byte) bool) {
		for i := range len(b) / 3 {
			if !yield(b[3*i : 3*i+3]) {
				return
			}
		}
	}
}
func main() {
	content, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatalf("failed to open input file: %v", err)
	}

	codonProtein := map[string]byte{
		"UUU": 'F',
		"CUU": 'L',
		"AUU": 'I',
		"GUU": 'V',
		"UUC": 'F',
		"CUC": 'L',
		"AUC": 'I',
		"GUC": 'V',
		"UUA": 'L',
		"CUA": 'L',
		"AUA": 'I',
		"GUA": 'V',
		"UUG": 'L',
		"CUG": 'L',
		"AUG": 'M',
		"GUG": 'V',
		"UCU": 'S',
		"CCU": 'P',
		"ACU": 'T',
		"GCU": 'A',
		"UCC": 'S',
		"CCC": 'P',
		"ACC": 'T',
		"GCC": 'A',
		"UCA": 'S',
		"CCA": 'P',
		"ACA": 'T',
		"GCA": 'A',
		"UCG": 'S',
		"CCG": 'P',
		"ACG": 'T',
		"GCG": 'A',
		"UAU": 'Y',
		"CAU": 'H',
		"AAU": 'N',
		"GAU": 'D',
		"UAC": 'Y',
		"CAC": 'H',
		"AAC": 'N',
		"GAC": 'D',
		"CAA": 'Q',
		"AAA": 'K',
		"GAA": 'E',
		"CAG": 'Q',
		"AAG": 'K',
		"GAG": 'E',
		"UGU": 'C',
		"CGU": 'R',
		"AGU": 'S',
		"GGU": 'G',
		"UGC": 'C',
		"CGC": 'R',
		"AGC": 'S',
		"GGC": 'G',
		"CGA": 'R',
		"AGA": 'R',
		"GGA": 'G',
		"UGG": 'W',
		"CGG": 'R',
		"AGG": 'R',
		"GGG": 'G',
	}

	var protein strings.Builder

	for c := range Codons(content) {
		codon := string(c)
		if codon == "UAA" || codon == "UAG" || codon == "UGA" {
			break
		}
		p, ok := codonProtein[codon]
		if !ok {
			break
		}

		protein.WriteByte(p)
	}
	fmt.Println(protein.String())
}
