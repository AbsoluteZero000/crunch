package huffman

import (
	"github.com/absolutezero000/encoding/huffman"
	"testing"
)

func checkIfEqual(root1 *huffman.Node, root2 *huffman.Node) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil || root2 == nil {
		return false
	}
	return root1.Char == root2.Char && root1.Freq == root2.Freq && checkIfEqual(root1.Left, root2.Left) && checkIfEqual(root1.Right, root2.Right)
}
func TestSeralization(t *testing.T) {
	testCases := []struct {
		name string
		root *huffman.Node
	}{
		{
			name: "simple tree with two characters",
			root: &huffman.Node{
				Char: 'a',
				Freq: 5,
				Left: &huffman.Node{
					Char:  'b',
					Freq:  2,
					Left:  nil,
					Right: nil,
				},
				Right: &huffman.Node{
					Char:  'c',
					Freq:  3,
					Left:  nil,
					Right: nil,
				},
			},
		},
		{
			name: "more complex tree",
			root: &huffman.Node{
				Char: 0,
				Freq: 10,
				Left: &huffman.Node{
					Char:  'd',
					Freq:  4,
					Left:  nil,
					Right: nil,
				},
				Right: &huffman.Node{
					Char: 0,
					Freq: 6,
					Left: &huffman.Node{
						Char:  'e',
						Freq:  3,
						Left:  nil,
						Right: nil,
					},
					Right: &huffman.Node{
						Char:  'f',
						Freq:  3,
						Left:  nil,
						Right: nil,
					},
				},
			},
		},
		{
			name: "tree with only root",
			root: &huffman.Node{
				Char:  'g',
				Freq:  7,
				Left:  nil,
				Right: nil,
			},
		},
		{
			name: "tree with deep nesting",
			root: &huffman.Node{
				Char: 0,
				Freq: 15,
				Left: &huffman.Node{
					Char:  'h',
					Freq:  5,
					Left:  nil,
					Right: nil,
				},
				Right: &huffman.Node{
					Char: 0,
					Freq: 10,
					Left: &huffman.Node{
						Char:  'i',
						Freq:  4,
						Left:  nil,
						Right: nil,
					},
					Right: &huffman.Node{
						Char:  'j',
						Freq:  6,
						Left:  nil,
						Right: nil,
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			serialized, err := huffman.SerializeTree(tc.root)
			if err != nil {
				t.Errorf("Deserialization failed")
			}
			root, err := huffman.DeserializeTree(serialized)
			if err != nil {
				t.Errorf("Deserialization failed")
			}
			if !checkIfEqual(tc.root, root) {
				t.Errorf("Deserialization failed")
			}
		})
	}

}
