/* Alice has some number of cards and she wants to rearrange the cards into groups so that each group is of size groupSize, and consists of groupSize consecutive cards.

Given an integer array hand where hand[i] is the value written on the ith card and an integer groupSize, return true if she can rearrange the cards, or false otherwise.

 

Example 1:

Input: hand = [1,2,3,6,2,3,4,7,8], groupSize = 3
Output: true
Explanation: Alice's hand can be rearranged as [1,2,3],[2,3,4],[6,7,8]
Example 2:

Input: hand = [1,2,3,4,5], groupSize = 4
Output: false
Explanation: Alice's hand can not be rearranged into groups of 4.

 

Constraints:

1 <= hand.length <= 104
0 <= hand[i] <= 109
1 <= groupSize <= hand.length
*/
package main

import (
	"fmt"
	"container/heap"
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

func (h *MinHeap) Top() int {
    return (*h)[0]
}

func isNStraightHand(hand []int, groupSize int) bool {
    if len(hand)%groupSize != 0 {
        return false
    }

    h := &MinHeap{}
    heap.Init(h) 

    for _,element := range hand {
        heap.Push(h, element)
    }

    for h.Len() != 0 {
        prev := -1
        var arr []int
        for i:=0; i<groupSize; i++ {
            if h.Len() != 0 {
                if prev == -1 || (h.Top() - prev == 1) {
                    prev = h.Top()
                } else {
                    i--
                    arr = append(arr, h.Top())
                }
                heap.Pop(h)
            }
        }

        for _,element := range arr {
            heap.Push(h, element)
        }

        if h.Len() % groupSize != 0 {
            return false
        }
    }

    return true
}

func main() {
	hand1 := []int{1,2,3,6,2,3,4,7,8}
	groupSize1 := 3
	fmt.Println(isNStraightHand(hand1, groupSize1)) // Output: true

	hand2 := []int{1,2,3,4,5}
	groupSize2 := 4
	fmt.Println(isNStraightHand(hand2, groupSize2)) // Output: false
}