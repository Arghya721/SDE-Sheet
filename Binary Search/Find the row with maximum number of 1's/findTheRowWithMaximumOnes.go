/*
	Given a boolean 2D array of n x m dimensions, consisting of only 1's and 0's, where each row is sorted. Find the 0-based index of the first row that has the maximum number of 1's. Return the 0-based index of the first row that has the most number of 1s. If no such row exists, return -1.

Examples :

Input: n = 4, m = 4, arr[][] = [[0, 1, 1, 1],[0, 0, 1, 1],[1, 1, 1, 1],[0, 0, 0, 0]]
Output: 2
Explanation: Row 2 contains 4 1's (0-based indexing).
Input: n = 2, m = 2, arr[][] = [[0, 0], [1, 1]]
Output: 1
Explanation: Row 1 contains 2 1's (0-based indexing).

Expected Time Complexity: O(n+m)
Expected Auxiliary Space: O(1)

Constraints:
1 ≤ n, m ≤ 103
0 ≤ arr[i][j] ≤ 1
*/
package main

func countOnes(arr [][]int, i int) int {
	var (
		low  = 0
		high = len(arr[i]) - 1
	)

	for low <= high {
		mid := (low + high) >> 1
		if arr[i][mid] == 1 {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return low
}

func rowWithMax1s(arr [][]int, n int, m int) int {
	var (
		ans     = -1
		prev    = 0
		maxOnes = 0
	)

	for i := 0; i < n; i++ {
		currRows := m - countOnes(arr, i)
		maxOnes = max(maxOnes, currRows)
		if prev != maxOnes {
			prev = maxOnes
			ans = i
		}
	}
	return ans
}


func main() {
	// Test case 1
	arr := [][]int{{0, 1, 1, 1}, {0, 0, 1, 1}, {1, 1, 1, 1}, {0, 0, 0, 0}}
	n := 4
	m := 4
	println(rowWithMax1s(arr, n, m)) // Output: 2

	// Test case 2
	arr = [][]int{{0, 0}, {1, 1}}
	n = 2
	m = 2
	println(rowWithMax1s(arr, n, m)) // Output: 1

	arr = [][]int{}
	n = 0
	m = 0
	println(rowWithMax1s(arr, n, m)) // Output: -1
}