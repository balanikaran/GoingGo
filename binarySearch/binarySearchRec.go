package main

import (
	"fmt"
)

func binarySearchRecursive(array []int, toSearch int, low int, high int) int {

	if low > high {
		return -1
	}

	mid := int((low + high) / 2)

	if array[mid] == toSearch {
		return mid
	}

	if array[mid] > toSearch {
		return binarySearchRecursive(array, toSearch, low, mid-1)
	} else {
		return binarySearchRecursive(array, toSearch, mid+1, high)
	}

}

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	index := binarySearchRecursive(array, 12, 0, len(array)-1)

	fmt.Println("Result = ", index)
}
