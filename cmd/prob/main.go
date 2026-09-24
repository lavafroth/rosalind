package main

import (
	"bytes"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	content, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatalf("failed to open input file: %v", err)
	}

	var gcContent []float64
	strand, probBytes, ok := bytes.Cut(content, []byte{'\n'})
	if !ok {
		log.Fatalln("failed to find separate dna strand and input probabilities")
	}
	for prob := range bytes.SplitSeq(bytes.TrimSpace(probBytes), []byte{' '}) {
		p, err := strconv.ParseFloat(string(prob), 64)
		if err != nil {
			log.Fatalf("failed to parse probability value: %v", err)
		}

		gcContent = append(gcContent, p)
	}

	var probs []float64
	for _, gc := range gcContent {
		claimC := gc / 2.0
		claimA := 0.5 - claimC
		jointProb := 1.0
		for _, ch := range strand {
			if ch == 'A' || ch == 'T' {
				jointProb *= claimA
			}
			if ch == 'C' || ch == 'G' {
				jointProb *= claimC
			}
		}
		probs = append(probs, math.Log10(jointProb))
	}
	probString := fmt.Sprint(probs)
	probString = probString[1 : len(probString)-1]
	fmt.Println(probString)
}
