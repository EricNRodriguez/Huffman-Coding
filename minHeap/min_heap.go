package minHeap

import (
	"errors"
)

type MinHeap struct {
	Size int
	Nodes []*minHeapNode
}

type minHeapNode struct {
	character rune
	frequency int
}


func (h *MinHeap)Insert(char rune, freq int) rune {
	h.Nodes = append(h.Nodes, &minHeapNode{
		char,
		freq,
	})
	i := len(h.Nodes) - 1
	for i != 0 && (h.Nodes[i].frequency < h.Nodes[(i-1)/2].frequency) {
		node := h.Nodes[i]
		h.Nodes[i] = h.Nodes[(i-1)/2]
		h.Nodes[(i-1)/2] = node
		i = (i-1)/2
	}

	return char
}

func (h *MinHeap)Remove() (n *minHeapNode, err error) {
	if len(h.Nodes) == 0 {
		return nil, errors.New("[ERR] Empty heap")
	}
	n = h.Nodes[0]
	h.Nodes[0] = h.Nodes[len(h.Nodes)-1] //make root equal last inserted element
	h.Nodes = h.Nodes[:len(h.Nodes)-1] //remove last element
	//down heap
	i := 0
	for 2*i + 1 < len(h.Nodes)-1  || 2*i + 2 < len(h.Nodes)-1 {
		if 2*i + 1 < len(h.Nodes)-1  && 2*i + 2 < len(h.Nodes)-1 {
			switch h.Nodes[2*i+1].frequency < h.Nodes[2*i+2].frequency {
			case true:
				left := h.Nodes[2*i+1]
				h.Nodes[2*i+1] = h.Nodes[i]
				h.Nodes[i] = left
				i = 2*i + 1
			case false:
				right := h.Nodes[2*i+2]
				h.Nodes[2*i+2] = h.Nodes[i]
				h.Nodes[i] = right
				i = 2*i + 2
			}
		} else if 2*i + 1 >= len(h.Nodes)-1 {
			right := h.Nodes[2*i + 2]
			h.Nodes[2*i + 2] = h.Nodes[i]
			h.Nodes[i] = right
			i = 2*i + 2
		} else {
			left := h.Nodes[2*i + 1]
			h.Nodes[2*i + 1] = h.Nodes[i]
			h.Nodes[i] = left
			i = 2*i + 1
		}
	}
	return
}


