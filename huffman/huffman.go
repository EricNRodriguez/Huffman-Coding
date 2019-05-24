package huffman

import (
	"Huffman/minHeap"
	"fmt"
)

type node struct {
	value rune
	encoding []int
	left *node
	right *node
}

func EncodeMessage(message string) (encodedMessage []int, encoding map[rune][]int, err error) {
	charFreq := getCharacterFrequencies([]rune(message))
	charFreqHeap := minHeap.MinHeap{}
	for char, freq := range charFreq {
		charFreqHeap.Insert(&node{value:char}, freq)
	}
	for charFreqHeap.Size > 1 {
		minOne, _ := charFreqHeap.RemoveMin()
		minTwo, _ := charFreqHeap.RemoveMin()
		charFreqHeap.Insert(
			&node{
				69,
				[]int{},
				minOne.Value.(*node),
				minTwo.Value.(*node),
			},
			minOne.Frequency + minTwo.Frequency)
	}
	huffmanTreeRoot, _ := charFreqHeap.RemoveMin()
	encoding = make(map[rune][]int)
	traverseAndRecord(huffmanTreeRoot.Value.(*node), encoding)


	for _, char := range []rune(message) {
		for _, b := range encoding[char] {
			encodedMessage = append(encodedMessage, b)
		}
	}
	return
}

func getCharacterFrequencies(chars []rune) (frequencies map[rune]int) {
	frequencies = make(map[rune]int)
	for _, char := range chars {
		if _, ok := frequencies[char]; ok {
			frequencies[char] += 1
		} else {
			frequencies[char] = 1
		}
	}
	return
}

//in order traversal, records path to every node
func traverseAndRecord(n *node, encoding map[rune][]int) {
	if n.left == nil && n.right == nil {
		encoding[n.value] = n.encoding
		fmt.Println(string(n.value))
	}
	if n.left != nil {
		n.left.encoding = append(n.encoding, 1)
		traverseAndRecord(n.left, encoding)
	}
	if n.right != nil {
		n.right.encoding = append(n.encoding, 0)
		traverseAndRecord(n.right, encoding)
	}
	return
}


