package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	// fmt.Println(sum([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
	// fmt.Println(countElement([]int{1, 2, 3, 4, 5}))
	// fmt.Println(greateCommonDivisor(270, 192))
	fmt.Println(quickSort([]int{1, 6, 3, 6, 98, 1, 3, 6}))
}

func quickSort(nums []int) []int {

	// basic case
	if len(nums) < 2 {
		return nums
	}

	// pivot index
	pivotIndex := rand.IntN(len(nums))
	pivot := nums[pivotIndex]

	// empty slices
	var less []int
	var greater []int

	for i, num := range nums {
		// skip pivot element
		if i == pivotIndex {
			continue
		}

		if num <= pivot {
			less = append(less, num)
		} else {
			greater = append(greater, num)
		}
	}

	// recursion and combine
	result := append(quickSort(less), pivot)
	result = append(result, quickSort(greater)...)
	return result
}

func sum(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	return nums[0] + sum(nums[1:])
}

func countElement(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	return 1 + countElement(nums[1:])
}

func greateCommonDivisor(a, b int) int {
	if a == 0 {
		return b
	}

	if b == 0 {
		return a
	}

	return greateCommonDivisor(b, a%b)
}
