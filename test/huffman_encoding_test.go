package test

import (
	"Huffman/huffman"
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"
	"github.com/stretchr/testify/assert"
)

func TestRandomMessage(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	length := rand.Intn(100)
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	encodedMessage, huffmanTree, _ :=  huffman.EncodeMessage(b.String())
	message, err := huffman.DecodeMessage(encodedMessage, huffmanTree)
	if err != nil {
		fmt.Println(err)
	}
	assert.Equal(t, encodedMessage, message, "Message should be identical after encoding-decoding")
}