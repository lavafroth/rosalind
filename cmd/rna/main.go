package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	content, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatalf("failed to open input file: %v", err)
	}

	contents := strings.TrimSpace(string(content))
	fmt.Println(strings.ReplaceAll(contents, "T", "U"))

}
