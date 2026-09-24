package main

import (
	"bytes"
	"fmt"
	"github.com/lavafroth/rosalind/codon"
	"github.com/lavafroth/rosalind/revcomp"
	"log"
	"os"
	"strings"
)

var codonProtein = map[string]byte{
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

func transcribe(content []byte, set map[string]struct{}) {
	var protein strings.Builder
	started := false

	for c := range codon.Chunks(content) {
		codon := string(c)
		if !started && codon != "ATG" {
			continue
		}
		started = true

		if codon == "TAA" || codon == "TAG" || codon == "TGA" {
			started = false
			proteinString := protein.String()
			for i, prot := range proteinString {
				if prot == 'M' {
					set[proteinString[i:]] = struct{}{}
				}
			}
			protein.Reset()
			continue
		}

		p, ok := codonProtein[codon]
		if !ok {
			break
		}
		protein.WriteByte(p)
	}
}

func main() {
	content, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatalf("failed to open input file: %v", err)
	}

	var sequence []byte
	for line := range bytes.Lines(content) {
		line = bytes.TrimSpace(line)
		if len(line) == 0 || line[0] == '>' {
			continue
		}
		sequence = append(sequence, line...)
	}

	revComp, err := revcomp.ReverseComplement(sequence)
	if err != nil {
		log.Fatal(err)
	}

	proteins := make(map[string]struct{})

	referenceFrames := [][]byte{sequence, sequence[1:], sequence[2:], revComp, revComp[1:], revComp[2:]}
	for _, referenceFrame := range referenceFrames {
		transcribe(referenceFrame, proteins)
	}
	for protein := range proteins {
		fmt.Println(protein)
	}
}
