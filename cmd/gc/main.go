package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {
	contents, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var maxId []byte
	maxGc := 0.0

	for chunk := range bytes.SplitSeq(contents, []byte(">")) {
		if len(chunk) == 0 {
			continue
		}

		id, sequence, found := bytes.Cut(chunk, []byte("\n"))
		if !found {
			break
		}

		gcCount, blank := 0, 0

		for _, char := range sequence {
			if char == 'C' || char == 'G' {
				gcCount += 1
			}
			if char == '\n' {
				blank += 1
			}
		}

		gc := float64(gcCount) * 100 / float64(len(sequence)-blank)
		if gc > maxGc {
			maxGc = gc
			maxId = id
		}
	}

	fmt.Println(string(maxId))
	fmt.Println(maxGc)
}
