/* You are given an m x n integer matrix matrix with the following two properties:

Each row is sorted in non-decreasing order.
The first integer of each row is greater than the last integer of the previous row.
Given an integer target, return true if target is in matrix or false otherwise.

You must write a solution in O(log(m * n)) time complexity.

 

Example 1:


Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
Output: true
Example 2:


Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
Output: false
 

Constraints:

m == matrix.length
n == matrix[i].length
1 <= m, n <= 100
-104 <= matrix[i][j], target <= 104
*/
package main

func searchMatrix(matrix [][]int, target int) bool {
    var (
        low = 0
        high = len(matrix) * len(matrix[0]) - 1
    )

    for low <= high {
        mid := (low+high)/2
        if matrix[mid/len(matrix[0])][mid%len(matrix[0])] == target {
            return true
        }

        if matrix[mid/len(matrix[0])][mid%len(matrix[0])] < target {
            low = mid + 1
        } else {
            high = mid - 1
        }
    }
    return false
}

func main() {
	// Test case 1
	matrix := [][]int{{1,3,5,7},{10,11,16,20},{23,30,34,60}}
	target := 3
	println(searchMatrix(matrix, target)) // Output: true

	// Test case 2
	matrix = [][]int{{1,3,5,7},{10,11,16,20},{23,30,34,60}}
	target = 13
	println(searchMatrix(matrix, target)) // Output: false
}