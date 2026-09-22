package revcomp

import (
	"fmt"
)

func ReverseComplement(s []byte) ([]byte, error) {
	comp := map[byte]byte{'A': 'T', 'T': 'A', 'C': 'G', 'G': 'C'}
	revComp := make([]byte, len(s))
	for i := range len(s) {
		n := s[len(s)-i-1]
		r, ok := comp[n]
		if !ok {
			return nil, fmt.Errorf("found non nucleotide character in dna sequence: 0x%08x", n)
		}
		revComp[i] = r
	}
	return revComp, nil
}
