package main

import (
	"fmt"
)

func main() {
	target := []int{4, 2, 1, 1, 9, 8, 8}
	size := 7

	fmt.Println(target)
	quickSort(target, 0, size)
	fmt.Println(target)
}

func quickSort(target []int, low, high int) {
	if high-low <= 1 {
		return
	}

	pivot_value := target[low]
	left, right := low, high

	for i := low; i < high; i++ {
		if target[i] < pivot_value {
			target[i], target[left] = target[left], target[i]
			left++
		}

		if target[i] > pivot_value {
			right--
			target[i], target[right] = target[right], target[i]
		}
	}
	quickSort(target, low, left)
	quickSort(target, right, high)
}
