/* Given an array arr[] and an integer k where k is smaller than the size of the array, the task is to find the kth smallest element in the given array. It is given that all array elements are distinct.

Note:-  l and r denotes the starting and ending index of the array.

Example 1:

Input:
n = 6
arr[] = 7 10 4 3 20 15
k = 3
l=0 r=5

Output : 
7

Explanation :
3rd smallest element in the given 
array is 7.
Example 2:

Input:
n = 5
arr[] = 7 10 4 20 15
k = 4 
l=0 r=4

Output : 
15

Explanation :
4th smallest element in the given 
array is 15.
Your Task:
You don't have to read input or print anything. Your task is to complete the function kthSmallest() which takes the array arr[], integers l and r denoting the starting and ending index of the array and an integer k as input and returns the kth smallest element.
 
Expected Time Complexity: O(n*log(n) )
Expected Auxiliary Space: O(log(n))
Constraints:
1 <= n <= 105
l = 0
r = N-1
1<= arr[i] <= 105
1 <= k <= n

*/
package main

import (
	"container/heap"
	"fmt"
)

// MinHeap struct
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

func kthSmallest(arr []int, l, r, k int) int {
	h := &MinHeap{}
	heap.Init(h)

	// Push elements from arr[l] to arr[r] into the heap
	for i := l; i <= r; i++ {
		heap.Push(h, arr[i])
	}

	// Pop the smallest element k-1 times
	for k > 1 {
		heap.Pop(h)
		k--
	}

	// The root of the heap is the k-th smallest element
	return heap.Pop(h).(int)
}

func main() {
	arr1 := []int{12, 3, 5, 7, 19}
	k1 := 2
	fmt.Println(kthSmallest(arr1, 0, len(arr1)-1, k1)) // Output: 5

	arr2 := []int{7, 10, 4, 3, 20, 15}
	k2 := 3
	fmt.Println(kthSmallest(arr2, 0, len(arr2)-1, k2)) // Output: 7
}
