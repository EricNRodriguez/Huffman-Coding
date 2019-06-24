package bitStack

import (
	"errors"
)

type BitStack struct {
	Stack 	[]uint64
	Height 	int
}

func (b *BitStack) Push(bit uint64) (err error) {
	if bit != 0 && bit != 1 {
		err = errors.New("[ERR] Invalid bit provided ")
	}
	if b.Height % 64 == 0 {
		b.Stack = append(b.Stack, bit)
		b.Height++
		return
	}
	b.Stack[len(b.Stack)-1] = b.Stack[len(b.Stack)-1] << 1
	if bit == 1 {
		b.Stack[len(b.Stack)-1] |=  1
	}
	b.Height++
	return
}

func (b *BitStack) Pop() (bit uint64, err error) {
	if b.Height == 0 {
		err = errors.New("[ERR] Empty stack ")
		return
	}
	if b.Stack[len(b.Stack)-1] % 2 != 0 {
		bit = 1
	}
	if b.Height % 64 == 1 {
		b.Stack = b.Stack[:len(b.Stack)-1]
	} else {
		b.Stack[len(b.Stack)-1] = b.Stack[len(b.Stack)-1] >> 1
	}
	b.Height--
	return
}




