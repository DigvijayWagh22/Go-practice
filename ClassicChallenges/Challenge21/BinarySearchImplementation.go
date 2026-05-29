package main

import "fmt"

func BinarySearch(arr []int, target int) int {
	n := len(arr)
	if n == 0 {
		return -1
	}

	low := 0
	high := n - 1

	for low <= high {
		mid := (low + high) / 2

		if arr[mid] == target {
			return mid
		} else if arr[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

func BinarySearchRecursive(arr []int, target int, left int, right int) int {
	if len(arr) == 0 || left > right {
		return -1
	}
	mid := (left + right) / 2
	if arr[mid] == target {
		return mid
	} else if arr[mid] > target {
		return BinarySearchRecursive(arr, target, left, mid-1)
	} else {
		return BinarySearchRecursive(arr, target, mid+1, right)
	}

}

func FindInsertPosition(arr []int, target int) int {

	n := len(arr)
	if n == 0 {
		return 0
	}
	low := 0
	high := n - 1
	mid := (low + high) / 2

	for low <= high {
		mid = (low + high) / 2
		if arr[mid] == target {
			return mid
		} else if arr[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return low
}

func main() {
	arr := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
	target := 7
	index := BinarySearch(arr, target)
	fmt.Printf("BinarySearch: %d found at index %d \n", target, index)

	recursiveIndex := BinarySearchRecursive(arr, target, 0, len(arr)-1)
	fmt.Printf("BinarySearchRecursive: %d found at index %d \n", target, recursiveIndex)

	insertTarget := 8
	insertPos := FindInsertPosition(arr, insertTarget)
	fmt.Printf("FindInsertPosition: %d should be inserted at index %d\n", insertTarget, insertPos)
}
