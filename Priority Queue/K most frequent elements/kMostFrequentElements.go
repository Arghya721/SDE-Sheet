/*
	Given an integer array nums and an integer k, return the k most frequent elements. You may return the answer in any order.

Example 1:

Input: nums = [1,1,1,2,2,3], k = 2
Output: [1,2]
Example 2:

Input: nums = [1], k = 1
Output: [1]

Constraints:

1 <= nums.length <= 105
-104 <= nums[i] <= 104
k is in the range [1, the number of unique elements in the array].
It is guaranteed that the answer is unique.

Follow up: Your algorithm's time complexity must be better than O(n log n), where n is the array's size.
*/
package main

import (
	"container/heap"
	"fmt"
)

// Element holds the value and its frequency
type Element struct {
	value     int
	frequency int
}

// MaxHeap struct
type MaxHeap []Element

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i].frequency > h[j].frequency } // Max-heap based on frequency
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(Element))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func topKFrequent(nums []int, k int) []int {
	h := &MaxHeap{}
	heap.Init(h)

	ump := make(map[int]int)

	for _, num := range nums {
		ump[num]++
	}

	for value, freq := range ump {
		heap.Push(h, Element{value, freq})
	}

	var ans []int

	for i := 0; i < k; i++ {
		top := heap.Pop(h).(Element)
		ans = append(ans, top.value)
	}

	return ans
}

func main() {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2)) // [1,2]
	fmt.Println(topKFrequent([]int{1}, 1))                // [1]
}
