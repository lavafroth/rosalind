package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	content, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatalf("failed to open input file: %v", err)
	}

	var n, k uint64
	_, err = fmt.Sscanf(string(content), "%d %d", &n, &k)
	if err != nil {
		log.Fatalf("failed to extract final generation number and lifespan months: %v", err)
	}

	if n < 3 {
		fmt.Println("1")
		return
	}

	var f1, f2 uint64 = 1, 1

	for range n - 2 {
		offsprings := f1*k + f2
		f1 = f2
		f2 = offsprings

	}
	fmt.Println(f2)

}
