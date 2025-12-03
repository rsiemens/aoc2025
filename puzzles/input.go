package puzzles

import (
	"io"
	"os"
	"log"
)

func loadInput(fpath string) string {
	file, err := os.Open(fpath)
	if err != nil {
		log.Fatalf("Failed to open input file")
	}

	input, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Failed to open input file")
	}

	return string(input)
}
