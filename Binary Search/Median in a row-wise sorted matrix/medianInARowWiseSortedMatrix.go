/* Given a row wise sorted matrix of size R*C where R and C are always odd, find the median of the matrix.

Example 1:

Input:
R = 3, C = 3
M = [[1, 3, 5], 
     [2, 6, 9], 
     [3, 6, 9]]
Output: 5
Explanation: Sorting matrix elements gives 
us {1,2,3,3,5,6,6,9,9}. Hence, 5 is median. 
 

Example 2:

Input:
R = 3, C = 1
M = [[1], [2], [3]]
Output: 2
Explanation: Sorting matrix elements gives 
us {1,2,3}. Hence, 2 is median.

Your Task:  
You don't need to read input or print anything. Your task is to complete the function median() which takes the integers R and C along with the 2D matrix as input parameters and returns the median of the matrix.

Expected Time Complexity: O(32 * R * log(C))
Expected Auxiliary Space: O(1)


Constraints:
1 <= R, C <= 400
1 <= matrix[i][j] <= 2000
*/
package main
import (
	"math"
)

func upperBound(arr []int, target, n int) int {
	low, high := 0, n
	for low < high {
		mid := (low+high)/2
		if arr[mid] <= target {
			low = mid+1
		} else {
			high = mid
		}
	}
	return low
}

func calculateSmallerThanMid(matrix [][]int, mid, r, c int) int {
	count := 0
	for i := 0; i < r; i++ {
		count += upperBound(matrix[i], mid, c)
	}
	return count
}

func median(matrix [][]int, r , c int) int {
	var (
		low = math.MaxInt
		high = math.MinInt
	)

	for i := 0; i < r; i++ {
		low = min(low, matrix[i][0])
		high = max(high, matrix[i][c-1])
	}

	req := (r*c)/2

	for low <= high {
		mid := (low+high)/2
		smallerThanMid := calculateSmallerThanMid(matrix, mid, r, c)
		if smallerThanMid <= req {
			low = mid+1
		} else {
			high = mid-1
		}
	}

	return low
}

func main() {
	matrix := [][]int{{1, 3, 5}, {2, 6, 9}, {3, 6, 9}}
	r, c := 3, 3
	println(median(matrix, r, c))

	matrix = [][]int{{1}, {2}, {3}}
	r, c = 3, 1
	println(median(matrix, r, c))	
}