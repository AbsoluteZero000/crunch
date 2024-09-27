package main

import (
	"fmt"
	"os"
	"sort"
)

type Pair struct {
	Char  rune
	Count int
}

func getSortedFrequency(s string) []Pair {
	frequencyMap := make(map[rune]int)

	for _, char := range s {
		frequencyMap[char]++
	}

	pairs := make([]Pair, 0, len(frequencyMap))
	for char, count := range frequencyMap {
		pairs = append(pairs, Pair{Char: char, Count: count})
	}

	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].Count == pairs[j].Count {
			return pairs[i].Char < pairs[j].Char
		}
		return pairs[i].Count > pairs[j].Count
	})

	return pairs
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ", os.Args[0], "filename")
		fmt.Println("No file specified")
		os.Exit(1)
	}

	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(data))

	sorted_slice := getSortedFrequency(string(data))

	fmt.Println(sorted_slice)

	for _, pair := range sorted_slice {
		fmt.Printf("Character: '%c', Frequency: %d\n", pair.Char, pair.Count)
	}

}
