/* Given k sorted arrays arranged in the form of a matrix of size k * k. The task is to merge them into one sorted array. Return the merged sorted array ( as a pointer to the merged sorted arrays in cpp, as an ArrayList in java, and list in python).


Examples :

Input: k = 3, arr[][] = {{1,2,3},{4,5,6},{7,8,9}}
Output: 1 2 3 4 5 6 7 8 9
Explanation: Above test case has 3 sorted arrays of size 3, 3, 3 arr[][] = [[1, 2, 3],[4, 5, 6],[7, 8, 9]]. The merged list will be [1, 2, 3, 4, 5, 6, 7, 8, 9].
Input: k = 4, arr[][]={{1,2,3,4},{2,2,3,4},{5,5,6,6},{7,8,9,9}}
Output: 1 2 2 2 3 3 4 4 5 5 6 6 7 8 9 9 
Explanation: Above test case has 4 sorted arrays of size 4, 4, 4, 4 arr[][] = [[1, 2, 2, 2], [3, 3, 4, 4], [5, 5, 6, 6], [7, 8, 9, 9 ]]. The merged list will be [1, 2, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 8, 9, 9].
Expected Time Complexity: O(k2*Log(k))
Expected Auxiliary Space: O(k2)

Constraints:
1 <= k <= 100
*/
package main

import (
	"container/heap"
	"fmt"
)

// MinHeap struct for a priority queue
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] } // Min-heap
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func mergeKArrays(arr [][]int) []int {
	h := &MinHeap{}
	heap.Init(h)

	// Push all elements from all arrays into the heap
	for _, subArr := range arr {
		for _, elem := range subArr {
			heap.Push(h, elem)
		}
	}

	// Pop elements from the heap and add to the result array
	result := make([]int, 0)
	for h.Len() > 0 {
		result = append(result, heap.Pop(h).(int))
	}

	return result
}

func main() {
	arr := [][]int{
		{1, 3, 5, 7},
		{2, 4, 6, 8},
		{0, 9, 10, 11},
	}
	result := mergeKArrays(arr)
	fmt.Println(result) // Output: [0 1 2 3 4 5 6 7 8 9 10 11]
}
