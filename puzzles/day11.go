package puzzles

import (
	"fmt"
	"strings"
)

type deviceGraph map[string][]string


func cacheKey(k string, dac, fft bool) string {
	key := k
	if dac {
		key += "_dac"
	}
	if fft {
		key += "_fft"
	}

	return key
}

func dfs(graph deviceGraph, k, target string, cache map[string]int, dac, fft bool) int {
	paths := graph[k]
	count := 0

	if k == "dac" {
		dac = true
	} else if k == "fft" {
		fft = true
	}

	if k == target && dac && fft {
		return 1
	}

	cacheK := cacheKey(k, dac, fft)
	if val, ok := cache[cacheK]; ok {
		return val
	}

	for _, next := range paths {
		count += dfs(graph, next, target, cache, dac, fft)
	}

	cache[cacheK] = count
	return count
}


func parseInput(input string) deviceGraph {
	graph := make(deviceGraph)
	for line := range strings.SplitSeq(strings.TrimSpace(input), "\n") {
		var key string
		var conns []string

		for i, v := range strings.Fields(line) {
			if i == 0 {
				key = strings.TrimSuffix(v, ":")
			} else {
				conns = append(conns, v)
			}
		}
		graph[key] = conns
	}
	return graph
}

func Day11() {
	input := loadInput("day11.txt")
	graph := parseInput(input)
	cache := make(map[string]int)
	fmt.Println("Day 11", dfs(graph, "svr", "out", cache, false, false))
}
