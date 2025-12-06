package puzzles

import (
	"io"
	"os"
	"log"
)

func loadInput(fname string) string {
	file, err := os.Open("inputs/" + fname)
	if err != nil {
		log.Fatalf("Failed to open input file")
	}

	input, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Failed to open input file")
	}

	return string(input)
}
