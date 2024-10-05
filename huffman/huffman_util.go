package huffman

import (
	"bytes"
	"container/heap"
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"math"
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

func bitStringToBytes(s string) []byte {
	numBytes := int(math.Ceil(float64(len(s)) / 8.0))
	result := make([]byte, numBytes)

	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			byteIndex := i / 8
			bitIndex := 7 - (i % 8)
			result[byteIndex] |= 1 << bitIndex
		}
	}

	return result
}

func bytesToBitString(data []byte, bitLength int) string {
	var result bytes.Buffer

	for i := 0; i < bitLength; i++ {
		byteIndex := i / 8
		bitIndex := 7 - (i % 8)

		if data[byteIndex]&(1<<bitIndex) != 0 {
			result.WriteByte('1')
		} else {
			result.WriteByte('0')
		}
	}

	return result.String()
}

func printCompressionStats(content string, codes map[rune]string, bitString string, compressionRatio float64) {
	colors := struct {
		primary   string
		secondary string
		accent    string
		text      string
		border    string
	}{
		primary:   "#7D56F4",
		secondary: "#2D3748",
		accent:    "#48BB78",
		text:      "#FAFAFA",
		border:    "#4A5568",
	}

	maxWidth := 100

	baseStyle := lipgloss.NewStyle().
		PaddingLeft(1).
		PaddingRight(1).
		BorderStyle(lipgloss.RoundedBorder())

	headerStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color(colors.text)).
		Background(lipgloss.Color(colors.primary)).
		PaddingLeft(2).
		PaddingRight(2).
		MarginBottom(1)

	contentBoxStyle := baseStyle.
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color(colors.border)).
		Padding(1).
		Width(maxWidth)

	codeRowStyle := lipgloss.NewStyle().
		PaddingLeft(1).
		PaddingRight(1).
		MarginBottom(1).
		Width(maxWidth - 4)

	statStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color(colors.accent))

	fmt.Println(headerStyle.Render("Original Content"))
	contentChunks := splitLongString(content, maxWidth-6)
	formattedContent := lipgloss.JoinVertical(lipgloss.Left, contentChunks...)
	fmt.Println(contentBoxStyle.Render(formattedContent))
	fmt.Println()

	fmt.Println(headerStyle.Render("Huffman Codes"))
	codeBox := ""
	for char, code := range codes {
		var row string
		if char == '\n' {
			row = fmt.Sprintf("Character: '\\n' │ Code: %-20s", code)
		} else {
			row = fmt.Sprintf("Character: '%c'  │ Code: %-20s", char, code)
		}
		codeBox += codeRowStyle.Render(row) + "\n"
	}
	fmt.Println(contentBoxStyle.Render(codeBox))
	fmt.Println()

	fmt.Println(headerStyle.Render("Bit String"))
	bitStringChunks := splitLongString(bitString, maxWidth-6)
	formattedBitString := lipgloss.JoinVertical(lipgloss.Left, bitStringChunks...)
	fmt.Println(contentBoxStyle.Render(formattedBitString))
	fmt.Println()

	statsStyle := contentBoxStyle.
		Width(50).
		Align(lipgloss.Center)

	stats := fmt.Sprintf(
		"%s: %d bits\n%s: %d bits\n%s: %.2f%%",
		statStyle.Render("Original size"),
		len(content)*8,
		statStyle.Render("Compressed size"),
		len(bitString),
		statStyle.Render("Compression ratio"),
		compressionRatio,
	)
	fmt.Println(headerStyle.Render("Compression Statistics"))
	fmt.Println(statsStyle.Render(stats))
}

func splitLongString(s string, chunkSize int) []string {
	var chunks []string
	runes := []rune(s)

	for i := 0; i < len(runes); i += chunkSize {
		end := i + chunkSize
		if end > len(runes) {
			end = len(runes)
		}
		chunks = append(chunks, string(runes[i:end]))
	}
	return chunks
}
