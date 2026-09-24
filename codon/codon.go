package codon

import "iter"

func Chunks(b []byte) iter.Seq[[]byte] {
	return func(yield func([]byte) bool) {
		for i := range len(b) / 3 {
			if !yield(b[3*i : 3*i+3]) {
				return
			}
		}
	}
}
