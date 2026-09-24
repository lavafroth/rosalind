package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"slices"
)

type Summary struct {
	name   string
	prefix string
	suffix string
}

func main() {
	content, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatalf("failed to open input file: %v", err)
	}

	Bases := []byte{'A', 'C', 'G', 'T'}
	var collection []Summary
	for entry := range bytes.SplitSeq(content, []byte{'>'}) {
		if len(entry) == 0 {
			continue
		}
		name, seq, ok := bytes.Cut(entry, []byte{'\n'})
		if !ok {
			log.Fatal("unable to split fasta entry into label and sequences")
		}

		var prefix []byte
		var suffix []byte
		for _, ch := range seq {
			if slices.Index(Bases, ch) != -1 {
				prefix = append(prefix, ch)
			}
			if len(prefix) == 3 {
				break
			}
		}
		for i := range len(seq) {
			ch := seq[len(seq)-1-i]
			if slices.Index(Bases, ch) != -1 {
				suffix = append(suffix, ch)
			}
			if len(suffix) == 3 {
				break
			}
		}
		slices.Reverse(suffix)
		collection = append(collection, Summary{string(name), string(prefix), string(suffix)})
	}

	for _, left := range collection {
		for _, right := range collection {
			if left.suffix == right.prefix && left.name != right.name {
				fmt.Println(left.name, right.name)
			}
		}
	}
}
