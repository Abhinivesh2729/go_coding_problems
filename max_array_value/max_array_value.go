//Find the Maximum Value in an Array

package main

import "fmt"

func main() {
	array := []int{1, 3, 2, 4, 7, 5, 9, 6, 8, 10}
	maxNum := findMaxnumber(&array)
	fmt.Printf("The max number is %v\n", maxNum)
}

func findMaxnumber(array *[]int) int {
	var maxNum int
	for _, number := range *array {
		if number > maxNum {
			maxNum = number
		}
	}
	return maxNum
}
