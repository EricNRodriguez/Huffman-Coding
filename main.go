package main

import (
	"Huffman/huffman"
	"fmt"
)

func main() {

	encodedMessage, huffmanTree, err := huffman.EncodeMessage("message to be encoded")
	message, err := huffman.DecodeMessage(encodedMessage, huffmanTree)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(encodedMessage.Stack, " here")
	fmt.Println(message)
}
