/* Given an integer N, print all the divisors of N in the ascending order.
 

Example 1:

Input : 20
Output: 1 2 4 5 10 20
Explanation: 20 is completely 
divisible by 1, 2, 4, 5, 10 and 20.

Example 2:

Input: 21191
Output: 1 21191
Explanation: As 21191 is a prime number,
it has only 2 factors(1 and the number itself).

Your Task:

Your task is to complete the function print_divisors() which takes N as input parameter and prints all the factors of N as space seperated integers in sorted order. You don't have to print any new line after printing the desired output. It will be handled in driver code.
 

Expected Time Complexity: O(sqrt(N))
Expected Space Complexity: O(sqrt(N))
 

Constraints:
1 <= N <= 105
*/
package main

import (
	"fmt"
)

func print_divisors(N int) {
	var (
		firstHalf []int
		secondHalf []int
	)

	for i := 1; i * i <= N; i++ {
		if N % i == 0 {
			firstHalf = append(firstHalf, i)
			if i != N / i {
				secondHalf = append(secondHalf, N / i)
			}
		}
	}

	for i := 0; i < len(firstHalf); i++ {
		fmt.Printf("%d ", firstHalf[i])
	}

	for i := len(secondHalf) - 1; i >= 0; i-- {
		fmt.Printf("%d ", secondHalf[i])
	}
}

func main() {
	print_divisors(20)
}