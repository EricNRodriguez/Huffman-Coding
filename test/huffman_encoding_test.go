package test

import (
	"Huffman/huffman"
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"strings"
	"testing"
	"time"
)

func TestRandomMessage(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	length := rand.Intn(100)
	b := strings.Builder{}
	for i := 0; i < length; i++ {
		b.WriteRune(rune(rand.Intn(126)))
	}
	encodedMessage, huffmanTree, err := huffman.EncodeMessage(b.String())
	if err != nil {
		assert.Error(t, err, "[ERR] Encoding message failed")
	}
	message, err := huffman.DecodeMessage(encodedMessage, huffmanTree)
	if err != nil {
		assert.Error(t, err, "[ERR] Decoding message failed")
	}
	assert.Equal(t, b.String(), message, "Message should be identical after encoding-decoding")
	fmt.Println(message)
}
