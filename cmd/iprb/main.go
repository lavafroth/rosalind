package main

import (
	"fmt"
	"log"
	"os"
)

func Choose2(n uint64) float64 {
	if n < 2 {
		return 0
	}
	m := float64(n)
	return m * (m - 1) / 2
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatalf("failed to open input file: %v", err)
	}
	var k, m, n uint64
	if _, err = fmt.Fscanf(file, "%d %d %d", &k, &m, &n); err != nil {
		log.Fatalf("failed to read counts of homozygous dominant, heterozygous, homozygous recessive factors from input file: %v", err)
	}

	pool := k + m + n
	total := Choose2(pool)
	kk := Choose2(k)
	km, kn := float64(k*m), float64(k*n)

	mm := 0.75 * Choose2(m)
	mn := 0.5 * float64(m*n)

	prob := (kk + km + kn + mm + mn) / total
	fmt.Println(prob)

}
