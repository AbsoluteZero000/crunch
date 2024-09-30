package main

import (
	"container/heap"
	"fmt"
	"os"
)

type Node struct {
	Char  rune
	Freq  int
	Left  *Node
	Right *Node
}

type PriorityQueue []*Node

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].Freq < pq[j].Freq }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Node)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

type Pair struct {
	Char  rune
	Count int
}

func getFrequency(s string) []Pair {
	frequencyMap := make(map[rune]int)
	for _, char := range s {
		frequencyMap[char]++
	}
	pairs := make([]Pair, 0, len(frequencyMap))
	for char, count := range frequencyMap {
		pairs = append(pairs, Pair{Char: char, Count: count})
	}

	return pairs
}

func buildHuffmanTree(pairs []Pair) *Node {
	pq := make(PriorityQueue, len(pairs))
	for i, pair := range pairs {
		pq[i] = &Node{Char: pair.Char, Freq: pair.Count}
	}
	heap.Init(&pq)

	for pq.Len() > 1 {
		left := heap.Pop(&pq).(*Node)
		right := heap.Pop(&pq).(*Node)
		newNode := &Node{
			Freq:  left.Freq + right.Freq,
			Left:  left,
			Right: right,
		}
		heap.Push(&pq, newNode)
	}

	return heap.Pop(&pq).(*Node)
}

func generateHuffmanCodes(root *Node, prefix string, codes map[rune]string) {
	if root == nil {
		return
	}

	if root.Char != 0 {
		codes[root.Char] = prefix
		return
	}

	generateHuffmanCodes(root.Left, prefix+"0", codes)
	generateHuffmanCodes(root.Right, prefix+"1", codes)
}

func decodeHelper(root *Node, encoded string, index int) (decodedChar string, newIndex int) {
	if root.Char != 0 {
		return string(root.Char), index
	}

	if encoded[index] == '0' {
		return decodeHelper(root.Left, encoded, index+1)
	} else {
		return decodeHelper(root.Right, encoded, index+1)
	}

}

func decodeHuffman(root *Node, encoded string, index int) string {
	decodedData := ""
	char := ""
	for index < len(encoded) {
		char, index = decodeHelper(root, encoded, index)
		decodedData += char
	}

	return decodedData
}

func encodeString(s string, codes map[rune]string) string {
	var encoded string
	for _, char := range s {
		encoded += codes[char]
	}
	return encoded
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
		os.Exit(1)
	}

	content := string(data)
	fmt.Println("Original content:")
	fmt.Println(content)

	Freq := getFrequency(content)
	fmt.Println("\nCharacter frequencies:")
	for _, pair := range Freq {
		fmt.Printf("Character: '%c', Frequency: %d\n", pair.Char, pair.Count)
	}

	root := buildHuffmanTree(Freq)
	codes := make(map[rune]string)
	generateHuffmanCodes(root, "", codes)

	fmt.Println("\nHuffman Codes:")
	for char, code := range codes {
		fmt.Printf("Character: '%c', Code: %s\n", char, code)
	}

	encoded := encodeString(content, codes)

	fmt.Println("\nEncoded string:")
	fmt.Println(encoded)

	fmt.Printf("\nOriginal size: %d bits\n", len(content)*8)
	fmt.Printf("Compressed size: %d bits\n", len(encoded))
	compressionRatio := float64(len(encoded)) / float64(len(content)*8) * 100
	fmt.Printf("Compression ratio: %.2f%%\n\n", compressionRatio)

	decodedData := decodeHuffman(root, encoded, 0)
	fmt.Printf("Decoded string: %s\n", decodedData)

}
