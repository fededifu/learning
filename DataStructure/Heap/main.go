package main

import "fmt"

// MaxHeap struct has a slice that holds the array
type MaxHeap struct {
	array []int
}

// Insert adds element to the heap
func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array) - 1)
}

// maxHeapifyUp will heapify from bottom top
func (h *MaxHeap) maxHeapifyUp(index int) {
	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

// maxHeapifyDown will heapify from top to bottom
func (h *MaxHeap) maxHeapifyDown(index int) {
	lastIndex := len(h.array) - 1
	l, r := left(index), right(index)
	childToCompare := 0

	// loop while index has at least one child
	for l <= lastIndex {
		if l == lastIndex { // when left child is the only child
			childToCompare = l
		} else if h.array[l] > h.array[r] { // when left child is larger
			childToCompare = l
		} else { // when right child is larger
			childToCompare = r
		}

		if h.array[index] < h.array[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		} else {
			return
		}
	}
}

// get the parent index
func parent(i int) int {
	return (i - 1) / 2
}

// get the left index
func left(i int) int {
	return 2*i + 1
}

// get the right index
func right(i int) int {
	return 2*i + 2
}

// swap keys in the array
func (h *MaxHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

// Extract returns the largest key and removes it from the heap
func (h *MaxHeap) Extract() int {
	extraced := h.array[0]
	l := len(h.array) - 1

	if len(h.array) == 0 {
		fmt.Println("cannot extract because array length is 0")
		return -1
	}

	h.array[0] = h.array[l]
	h.array = h.array[:l]

	h.maxHeapifyDown(0)

	return extraced
}

func main() {
	m := MaxHeap{}
	fmt.Println(m)

	builHeap := []int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144}
	for _, h := range builHeap {
		m.Insert(h)
		fmt.Println(m)
	}

	for i := 0; i < 5; i++ {
		m.Extract()
		fmt.Println(m)
	}
}
