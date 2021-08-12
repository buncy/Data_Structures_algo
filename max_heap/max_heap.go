package main

import "fmt"

//max heap struct has a slice that holds an array

type MaxHeap struct {
	array []int
}

//Insert adds an element to the heap

func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array) - 1)
}

//Extract returns the largest key, and removes it from the heap

func (h *MaxHeap) Extract() int {
	extracted := h.array[0]
	l := len(h.array) - 1

	//when array is empty
	if len(h.array) == 0 {
		fmt.Println("cannot extract, length of array is 0")
		return -1
	}

	//take out the last index and put it in root
	h.array[0] = h.array[l]
	h.array = h.array[:l]

	h.maxHeapifyDown(0)
	return extracted
}

//maxHeapifyUp will heapify from bottom up
func (h *MaxHeap) maxHeapifyUp(index int) {

	//if the parent index key is smaller than the currentIndex key
	for h.array[parent(index)] < h.array[index] {
		fmt.Println("===>\nthis is the current index : ", index, "\n======>")
		h.swap(parent(index), index)
		index = parent(index)
		fmt.Println("===>\nthis is the swap index : ", index, "\n======>")
	}

}

//maxHeapifyDown will heapify from top Down
func (h *MaxHeap) maxHeapifyDown(index int) {

	lastIndex := len(h.array) - 1
	l, r := leftChild(index), rightChild(index)
	childToCompare := 0

	//loop while index has at least one child
	for l <= lastIndex {
		//when left child is the only child
		if l == lastIndex {
			childToCompare = l
		} else if h.array[l] > h.array[r] { //when left child is larger than right
			childToCompare = l
		} else { //when right child is larger than left
			childToCompare = r
		}

		//compare array value of current index to larger child and swap if smaller
		if h.array[index] < h.array[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = leftChild(index), rightChild(index)
		} else {
			return
		}
	}

}

//get the parent index
func parent(i int) int {
	return (i - 1) / 2
}

//get the left child index
func leftChild(i int) int {
	return (i * 2) + 1
}

//get the right child index
func rightChild(i int) int {
	return (i * 2) + 2
}

//swap keys in the array
func (h *MaxHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

func main() {

	m := &MaxHeap{}
	fmt.Println(m)

	buildHeap := []int{10, 20, 30, 50, 5, 23, 45, 67, 56, 34, 55}

	for _, v := range buildHeap {
		m.Insert(v)
		fmt.Println(m)
	}

	for i := 0; i < 5; i++ {
		m.Extract()
		fmt.Println(m)
	}

	// index := m.array[parent(5)]

	// child := m.array[5]

	// fmt.Printf("the parent of %v is %v ", child, index)
}
