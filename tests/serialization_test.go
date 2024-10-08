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
				t.Errorf("Serialization failed")
			}
			root, err := huffman.DeserializeTree(serialized)
			if err != nil {
				t.Errorf("Deserialization failed")
			}
			if !checkIfEqual(tc.root, root) {
				t.Errorf("The tree generated from the deserialization is not equal to the original tree")
			}
		})
	}

}

func TestTreeDataSeralization(t *testing.T) {
	testCases := []struct {
		name  string
		input string
	}{
		{
			name:  "Simple test",
			input: "Hello, World!",
		},
		{
			name:  "Longer test",
			input: "Hello from NeoVim this is a long test to figure out if my huffman encoding function is working properly or not",
		},
		{
			name:  "Repeated characters",
			input: "aaaabbbcccdddeeee",
		},
		{
			name:  "really huge test",
			input: "aaaaaaaaljkfsd;laskdjflkdhfaklhsdfkjlasdghflkagsfkgdjfdsagklfgaskdljfgjkasdgfkjlasgklfagskjgflksaglkjgfasdklgslkfjgklbbbbbbccccccdddddeeeeeeffffffgggggghhhhhiiiiiijjjjjkkkkklllllmmmmnnoooopppppqqqqrrrrssssttttuuuuvvvwwwxxxxyyyyzzzz",
		},
		{
			name:  "HUGE TEST",
			input: "The sun was setting magnificently over the distant mountains, casting long shadows across the verdant valley below, where countless wildflowers swayed gently in the evening breeze that whispered through the ancient oak trees standing sentinel along the weathered stone walls that had protected this peaceful sanctuary for generations untold, their branches reaching skyward like nature's own cathedral spires pointing toward the increasingly starlit heavens above, while down in the meadow, a family of deer cautiously emerged from the dense forest undergrowth to graze upon the tender spring grass that carpeted the earth in a luxuriant emerald tapestry, their presence unnoticed by the busy songbirds who continued their melodious evening chorus as twilight slowly descended upon this tranquil scene, bringing with it the first flickering appearances of fireflies that danced like tiny lanterns among the deepening shadows, their ephemeral light creating a magical display that seemed to bridge the gap between day and night, earth and sky, reality and dreams, as the last rays of sunlight painted the clouds in brilliant hues of gold and purple and crimson that seemed to set the very heavens ablaze with celestial fire, while in the distance, the sound of a church bell tolled the hour with its resonant bronze voice that echoed across the peaceful countryside, calling the faithful to evening prayers just as it had done for centuries past, its familiar tones bringing comfort to all who heard them as they went about their evening routines in the scattered farmhouses and cottages that dotted the landscape like precious gems strewn across a velvet cloth, their windows beginning to glow with warm lamplight as families gathered together for their evening meals, sharing stories of their day's adventures and misadventures while outside the cycle of nature continued its eternal dance, as predictable as the tides yet ever-changing in its infinite variety, a testament to the boundless creativity of the natural world that surrounds us all in its magnificent embrace, offering both shelter and inspiration to those who take the time to observe its countless wonders with open eyes and grateful hearts, finding in its rhythms and patterns the very pulse of life itself, beating steadily onward through the ages like a cosmic drumbeat that echoes in the soul of every living thing, connecting us all in ways both seen and unseen to the great mystery that lies at the heart of existence itself, a mystery that has inspired poets and philosophers, scientists and seekers throughout the ages to ponder the deeper meanings that lie hidden beneath the surface of everyday reality, waiting to be discovered by those who dare to look beyond the obvious and delve into the profound questions that have challenged human understanding since the dawn of consciousness itself, when our ancient ancestors first gazed up at the star-filled sky and wondered about their place in the vast cosmic dance that continues to unfold around us with each passing moment of each passing day, as regular as clockwork yet as unpredictable as the flight of a butterfly carried on the summer wind.",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			encoded, err := huffman.SerializeData(tc.input, false)
			if err != nil {
				t.Errorf("Serialization failed")
			}

			decoded, err := huffman.DeserializeData(encoded, false)
			if err != nil {
				t.Errorf("Deserialization failed")
			}

			if decoded != tc.input {
				t.Errorf("Encoding/decoding mismatch. Original: %s, Got: %s", tc.input, decoded)
			}

			if decoded != tc.input {
				t.Errorf("Deserialization Failed")
			}
		})
	}
}
