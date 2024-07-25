/* Given a rod of length N inches and an array of prices, price[]. price[i] denotes the value of a piece of length i. Determine the maximum value obtainable by cutting up the rod and selling the pieces.

Note: Consider 1-based indexing.

Example 1:

Input:
N = 8
Price[] = {1, 5, 8, 9, 10, 17, 17, 20}
Output:
22
Explanation:
The maximum obtainable value is 22 by 
cutting in two pieces of lengths 2 and 
6, i.e., 5+17=22.
Example 2:

Input:
N=8
Price[] = {3, 5, 8, 9, 10, 17, 17, 20}
Output: 
24
Explanation: 
The maximum obtainable value is 
24 by cutting the rod into 8 pieces 
of length 1, i.e, 8*price[1]= 8*3 = 24. 
Your Task:  
You don't need to read input or print anything. Your task is to complete the function cutRod() which takes the array A[] and its size N as inputs and returns the maximum price obtainable.

Expected Time Complexity: O(N2)
Expected Auxiliary Space: O(N)

Constraints:
1 ≤ N ≤ 1000
1 ≤ Ai ≤ 105
*/
package main

import (
	"container/heap"
	"fmt"
)

// MaxHeap struct
type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] } // Max-heap
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func cutRod(price []int, n int) int {
	dp := make([]int, n+1)
	h := &MaxHeap{}
	heap.Init(h)

	for i := 1; i <= n; i++ {
		maxVal := price[i-1]
		for j := 1; j <= i/2; j++ {
			if maxVal < dp[j]+dp[i-j] {
				maxVal = dp[j] + dp[i-j]
			}
		}
		dp[i] = maxVal
		heap.Push(h, maxVal)
	}

	return (*h)[0] // Return the maximum value (top of the heap)
}

func main() {
	N := 8
	Price := []int{1, 5, 8, 9, 10, 17, 17, 20}

	result := cutRod(Price, N)
	fmt.Printf("Maximum obtainable value: %d\n", result)
}

