/*
	A peak element in a 2D grid is an element that is strictly greater than all of its adjacent neighbors to the left, right, top, and bottom.

Given a 0-indexed m x n matrix mat where no two adjacent cells are equal, find any peak element mat[i][j] and return the length 2 array [i,j].

You may assume that the entire matrix is surrounded by an outer perimeter with the value -1 in each cell.

You must write an algorithm that runs in O(m log(n)) or O(n log(m)) time.

Example 1:

Input: mat = [[1,4],[3,2]]
Output: [0,1]
Explanation: Both 3 and 4 are peak elements so [1,0] and [0,1] are both acceptable answers.
Example 2:

Input: mat = [[10,20,15],[21,30,14],[7,16,32]]
Output: [1,1]
Explanation: Both 30 and 32 are peak elements so [1,1] and [2,2] are both acceptable answers.

Constraints:

m == mat.length
n == mat[i].length
1 <= m, n <= 500
1 <= mat[i][j] <= 105
No two adjacent cells are equal.
*/
package main

import "fmt"

func calculateMaxElementIndex(mat [][]int, mid int) int {
	var (
		maxElement = -1
		maxIndex   = 0
	)

	for i := range mat {
		if mat[i][mid] > maxElement {
			maxElement = mat[i][mid]
			maxIndex = i
		}
	}
	return maxIndex
}

func findPeakGrid(mat [][]int) []int {
	var (
		low  = 0
		high = len(mat[0]) - 1
	)

	for low <= high {
		mid := (low + high) / 2
		maxElementIndex := calculateMaxElementIndex(mat, mid)
		var (
			left  = 0
			right = 0
		)

		if mid-1 >= 0 {
			left = mat[maxElementIndex][mid-1]
		} else {
			left = -1
		}

		if mid+1 < len(mat[0]) {
			right = mat[maxElementIndex][mid+1]
		} else {
			right = -1
		}

		if mat[maxElementIndex][mid] > left && mat[maxElementIndex][mid] > right {
			return []int{maxElementIndex, mid}
		}

		if mat[maxElementIndex][mid] < left {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return []int{}
}

func main() {
	fmt.Println(findPeakGrid([][]int{{1, 4}, {3, 2}}))                          // [0,1]
	fmt.Println(findPeakGrid([][]int{{10, 20, 15}, {21, 30, 14}, {7, 16, 32}})) // [1,1]
}
