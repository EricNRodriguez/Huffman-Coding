package minHeap

import (
	"errors"
)

type MinHeap struct {
	Size int
	Nodes []*MinHeapNode
}

type MinHeapNode struct {
	Value interface{}
	Frequency int
}

func (h *MinHeap) Insert(value interface{}, freq int) interface{} {
	h.Nodes = append(h.Nodes, &MinHeapNode{
		value,
		freq,
	})
	//up heap
	i := len(h.Nodes) - 1
	for i != 0 && (h.Nodes[i].Frequency < h.Nodes[(i-1)/2].Frequency) {
		node := h.Nodes[i]
		h.Nodes[i] = h.Nodes[(i-1)/2]
		h.Nodes[(i-1)/2] = node
		i = (i-1)/2
	}
	h.Size ++
	return value
}

func (h *MinHeap) RemoveMin() (n *MinHeapNode, err error) {
	if len(h.Nodes) == 0 {
		return nil, errors.New("[ERR] Empty heap")
	}
	n = h.Nodes[0] //root value of heap
	h.Nodes[0] = h.Nodes[len(h.Nodes)-1] //make root equal last inserted element
	h.Nodes = h.Nodes[:len(h.Nodes)-1] //remove last element
	//down heap
	i := 0
	for 2 * i + 1 <= len(h.Nodes) - 1  || 2 * i + 2 <= len(h.Nodes) - 1 {
		if 2 * i + 1 <= len(h.Nodes) - 1  && 2 * i + 2 <= len(h.Nodes) - 1 && h.Nodes[i].Frequency > h.Nodes[2*i+1].Frequency && h.Nodes[i].Frequency > h.Nodes[2*i+2].Frequency {
			if h.Nodes[2*i+1].Frequency < h.Nodes[2*i+2].Frequency { //two children
				left := h.Nodes[2*i+1]
				h.Nodes[2*i+1] = h.Nodes[i]
				h.Nodes[i] = left
				i = 2 * i + 1
			} else {
				right := h.Nodes[2*i+2]
				h.Nodes[2*i+2] = h.Nodes[i]
				h.Nodes[i] = right
				i = 2 * i + 2
			}
		} else if 2 * i + 1 > len(h.Nodes)-1 && h.Nodes[2 * i + 2].Frequency < h.Nodes[i].Frequency { // right child
			right := h.Nodes[2*i + 2]
			h.Nodes[2*i + 2] = h.Nodes[i]
			h.Nodes[i] = right
			i = 2*i + 2
		} else if  h.Nodes[2 * i + 1].Frequency < h.Nodes[i].Frequency { //left child
			left := h.Nodes[2*i + 1]
			h.Nodes[2*i + 1] = h.Nodes[i]
			h.Nodes[i] = left
			i = 2*i + 1
		} else {
			break
		}
	}
	h.Size --
	return
}


