package main

import "fmt"

func main() {
	nums := []int{5, 3, 6, 2, 10}
	fmt.Println(selectionSort(nums))
	fmt.Println(secondSelectionSort(nums))
	fmt.Println(selSort(nums))
}

func selectionSort(nums []int) []int {

	for i := 0; i < len(nums); i++ {
		smallInd := findSmall(nums[i:])
		nums[i], nums[smallInd+i] = nums[smallInd+i], nums[i]
	}

	return nums
}

func findSmall(nums []int) int {
	small := nums[0]
	smallIndex := 0

	for ind, val := range nums {
		if val < small {
			small = val
			smallIndex = ind
		}
	}

	return smallIndex
}

func secondSelectionSort(nums []int) []int {
	for i := 0; i < len(nums)-1; i++ {
		minIndex := i

		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[minIndex] {
				minIndex = j
			}
		}

		nums[i], nums[minIndex] = nums[minIndex], nums[i]
	}
	return nums
}

func selSort(nums []int) []int {
	for i := 0; i < len(nums)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[minIndex] {
				minIndex = j
			}
		}

		nums[i], nums[minIndex] = nums[minIndex], nums[i]
	}

	return nums
}
