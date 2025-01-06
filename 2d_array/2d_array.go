//sum of diagonals of 2d array

package main

import "fmt"

func main() {
	ar := [][]int32{{2, 3}, {5, 6}}

	result1, result2 := sumOfArrray(ar)
	fmt.Printf("D1 sum is %v\nD2 sum is %v\n", result1, result2)

}

func sumOfArrray(array [][]int32) (int32, int32) {
	var sum1 int32
	var sum2 int32

	sum1 = array[0][0] + array[1][1]
	sum2 = array[0][1] + array[1][0]

	return sum1, sum2
}
