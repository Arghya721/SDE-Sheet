package main

import (
	"fmt"
)

// MaxHeap struct
type MaxHeap struct {
	H []int
	s int
}

// NewMaxHeap initializes a new MaxHeap
func NewMaxHeap() *MaxHeap {
	return &MaxHeap{
		H: make([]int, 0, 10009),
		s: -1,
	}
}

// parent returns the parent node index of node i
func (heap *MaxHeap) parent(i int) int {
	return (i - 1) / 2
}

// leftChild returns the left child index of node i
func (heap *MaxHeap) leftChild(i int) int {
	return (2 * i) + 1
}

// rightChild returns the right child index of node i
func (heap *MaxHeap) rightChild(i int) int {
	return (2 * i) + 2
}

// shiftUp shifts the node up to maintain the heap property
func (heap *MaxHeap) shiftUp(i int) {
	for i > 0 && heap.H[heap.parent(i)] < heap.H[i] {
		heap.H[heap.parent(i)], heap.H[i] = heap.H[i], heap.H[heap.parent(i)]
		i = heap.parent(i)
	}
}

// shiftDown shifts the node down to maintain the heap property
func (heap *MaxHeap) shiftDown(i int) {
	maxIndex := i
	l := heap.leftChild(i)

	if l <= heap.s && heap.H[l] > heap.H[maxIndex] {
		maxIndex = l
	}
	r := heap.rightChild(i)

	if r <= heap.s && heap.H[r] > heap.H[maxIndex] {
		maxIndex = r
	}
	if i != maxIndex {
		heap.H[i], heap.H[maxIndex] = heap.H[maxIndex], heap.H[i]
		heap.shiftDown(maxIndex)
	}
}

// insert inserts a new element into the heap
func (heap *MaxHeap) insert(p int) {
	heap.s++
	if heap.s < len(heap.H) {
		heap.H[heap.s] = p
	} else {
		heap.H = append(heap.H, p)
	}
	heap.shiftUp(heap.s)
}

// extractMax extracts the maximum element from the heap
func (heap *MaxHeap) extractMax() int {
	if heap.s < 0 {
		panic("Heap is empty")
	}
	ans := heap.H[0]
	heap.H[0], heap.H[heap.s] = heap.H[heap.s], heap.H[0]
	heap.s--
	heap.shiftDown(0)
	return ans
}

func main() {
	heap := NewMaxHeap()
	heap.insert(3)
	heap.insert(2)
	heap.insert(15)
	heap.insert(5)
	heap.insert(4)
	heap.insert(45)

	fmt.Println(heap.extractMax()) // 45
	fmt.Println(heap.extractMax()) // 15
	fmt.Println(heap.extractMax()) // 5
}