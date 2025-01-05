//Find the Index of the First Occurrence of a Target Value

package main

import "fmt"

func main() {
	target := 5
	array := []int{1, 3, 5, 2, 5, 4}
	firstIndex := findFirst(&array, target)
	fmt.Printf("The first occurence of %v is at %v\n", target, firstIndex)
}

func findFirst(array *[]int, target int) int {

	for index, num := range *array {
		if num == target {
			return index
		}
	}
	return 0
}
