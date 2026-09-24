package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	content, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatalf("failed to open input file: %v", err)
	}

	alphabet, count, ok := bytes.Cut(content, []byte{'\n'})
	if !ok {
		log.Fatal("found no line delimiter between alphabet and count")
	}

	var letters []byte
	for _, x := range bytes.TrimSpace(alphabet) {
		if x == ' ' {
			continue
		}
		letters = append(letters, x)
	}

	slices.Sort(letters)

	n, err := strconv.Atoi(string(bytes.TrimSpace(count)))
	if err != nil {
		log.Fatal("failed to parse count of characters in each sequence")
	}

	genLexOrdered(n, letters)
}

func genLexOrdered(n int, letters []byte) {
	var slots = make([]int, n)
	tail := n - 1
	var s strings.Builder
	for {
		for _, v := range slots {
			s.WriteByte(letters[v])
		}
		fmt.Println(s.String())

		slots[tail] += 1
		for i := range slots {
			iRev := tail - i
			if slots[iRev] < len(letters) {
				break
			}

			if iRev <= 0 {
				return
			}
			slots[iRev-1]++
			slots[iRev] = 0
		}

		s.Reset()
	}
}
