//Find the Sum of All Elements in an Array

package main

import "fmt"

func main() {
	array := []int{1, 2, 3, 5}
	sum := sumOfArray(&array)
	fmt.Printf("The sum of the array is %v\n", sum)
}

func sumOfArray(array *[]int) int {
	var sum int
	for _, number := range *array {
		sum = sum + number
		fmt.Println(number)
	}
	return sum
}
