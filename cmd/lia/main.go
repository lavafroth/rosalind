package main

import (
	"fmt"
	"log"
	"math"
	"math/big"
	"os"
)

func main() {
	content, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatalf("failed to open input file: %v", err)
	}

	var gen, successes int64
	if _, err := fmt.Sscanf(string(content), "%d %d", &gen, &successes); err != nil {
		log.Fatalf("failed to parse generation number and required number of successful events: %v", err)
	}

	population := int64(1) << gen
	pSuccess := 0.25

	pUnsatisifed := 0.0
	for i := range successes {
		pUnsatisifed +=
			comb(population, i) * math.Pow(pSuccess, float64(i)) * math.Pow(1.0-pSuccess, float64(population-i))
	}

	fmt.Println(1.0 - pUnsatisifed)

}

func comb(n, k int64) float64 {
	if k > n {
		return 0
	}
	if k*2 > n {
		k = n - k
	}

	result := big.NewInt(1)
	for i := range k {
		j := i + 1
		result.Mul(result, big.NewInt(n-k+j))
		result.Div(result, big.NewInt(j))
	}
	value, _ := result.Float64()
	return value
}
