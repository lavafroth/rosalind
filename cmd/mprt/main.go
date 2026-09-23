package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func nGlycosylationIndices(s []byte) []int {
	window_size := 4
	iters := len(s) - window_size + 1
	var indices []int

	for start := range iters {
		first := s[start]
		second := s[start+1]
		third := s[start+2]
		fourth := s[start+3]
		if first == 'N' && second != 'P' && (third == 'S' || third == 'T') && fourth != 'P' {
			indices = append(indices, start+1)
		}
	}
	return indices
}

func parseFastaStrand(fasta []byte) []byte {
	strand := make([]byte, 0, len(fasta))
	for line := range bytes.Lines(fasta) {
		line = bytes.TrimSpace(line)
		if len(line) < 0 || line[0] == '>' {
			continue
		}

		strand = append(strand, line...)
	}
	return strand
}

func main() {
	content, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatalf("failed to open input file: %v", err)
	}

	for line := range bytes.Lines(content) {
		id, _, ok := bytes.Cut(line, []byte("_"))
		if !ok {
			id = line
		}
		id = bytes.TrimSpace(id)
		text := fetchProtein(id)
		strand := parseFastaStrand(text)
		indices := nGlycosylationIndices(strand)
		if len(indices) > 0 {
			fmt.Print(string(line))
			indicesStr := fmt.Sprint(indices)
			fmt.Println(indicesStr[1 : len(indicesStr)-1])
		}
	}

}

func fetchProtein(id []byte) []byte {
	url := fmt.Sprintf("https://rest.uniprot.org/uniprotkb/%s.fasta", string(id))
	log.Println("url:", url)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp.Status)
	defer resp.Body.Close()

	gz, err := gzip.NewReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer gz.Close()

	text, err := io.ReadAll(gz)
	if err != nil {
		log.Fatal(err)
	}

	return text
}
