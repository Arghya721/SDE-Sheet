/* Given an unsorted array, Arr[] of size N and that contains even number of occurrences for all numbers except two numbers. Find the two numbers in decreasing order which has odd occurrences.

Example 1:

Input:
N = 8
Arr = {4, 2, 4, 5, 2, 3, 3, 1}
Output: {5, 1} 
Explaination: 5 and 1 have odd occurrences.

Example 2:

Input:
N = 8
Arr = {1 7 5 7 5 4 7 4}
Output: {7, 1}
Explaination: 7 and 1 have odd occurrences.

Your Task:
You don't need to read input or print anything. Your task is to complete the function twoOddNum() which takes the array Arr[] and its size N as input parameters and returns the two numbers in decreasing order which have odd occurrences.


Expected Time Complexity: O(N)
Expected Auxiliary Space: O(1)


Constraints:
2 ≤ N ≤ 106
1 ≤ Arri ≤ 1012

*/
package main

import (
	"fmt"
)

func twoOddNum(arr []int) []int {
	xor := 0
	for _, element := range arr {
		xor ^= element
	}

	rightMostSetBit := xor ^ (xor&(xor-1))

	x := 0
	y := 0

	for _, element := range arr {
		if element&rightMostSetBit != 0 {
			x ^= element
		} else {
			y ^= element
		}
	}

	if x > y {
		return []int{x, y}
	}
	return []int{y, x}
}

func main() {
	fmt.Println(twoOddNum([]int{4, 2, 4, 5, 2, 3, 3, 1})) // [5, 1]
	fmt.Println(twoOddNum([]int{1, 7, 5, 7, 5, 4, 7, 4})) // [7, 1]
}