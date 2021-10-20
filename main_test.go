package main

import "testing"

func TestQuickSort1(t *testing.T) {
	target := []int{1, 2, 4}
	QuickSort(target)

	assertEquals(target, []int{1, 2, 4}, t)
}

func TestQuickSort2(t *testing.T) {
	target := []int{4, 2, 1}
	QuickSort(target)

	assertEquals(target, []int{1, 2, 4}, t)
}

func assertEquals(actual []int, expected []int, t *testing.T) {
	lenAc := len(actual)
	lenEx := len(expected)
	if lenAc != lenEx {
		t.Error("Slice sizes are not equal.", actual, expected)
	}

	for i, a := range actual {
		if a != expected[i] {
			t.Error("Different values ​​in slices.", actual, expected)
		}
	}
}
