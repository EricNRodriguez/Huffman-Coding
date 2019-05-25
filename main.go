package main

import (
	"Huffman-Coding/huffman"
	"fmt"
)

func main() {

	encodedMessage, huffmanTree, err := huffman.EncodeMessage("message that you would like to encode")
	message, err := huffman.DecodeMessage(encodedMessage, huffmanTree)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(encodedMessage)
	fmt.Println(message)
}
