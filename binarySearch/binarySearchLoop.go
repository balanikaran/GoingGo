package main

import (
	"fmt"
)

func binarySearchLoop(array []int, toSearch int, low int, high int) int {
	for low < high {
		mid := int((low + high) / 2)
		if array[mid] == toSearch {
			return mid
		} else if array[mid] > toSearch {
			high = mid
		} else {
			low = mid
		}
	}
	return -1
}

func main() {
	array := []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20}
	
	index := binarySearchLoop(array, 17, 0, len(array) - 1)

	fmt.Println("Result = ", index)
}
