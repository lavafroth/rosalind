package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatalf("failed to open input file: %v", err)
	}

	var vs float64
	for _, w := range []float64{1.0, 1.0, 1.0, 0.75, 0.5} {
		var v int
		fmt.Fscanf(file, "%d", &v)
		vs += w * float64(v)
	}
	fmt.Println(vs * 2)
}
