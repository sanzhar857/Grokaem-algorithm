package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 1, 2, 2, 2, 3, 4}
	inx := firstEntry(nums, 2)
	fmt.Println(inx)

	// numbers := []int{1, 3, 5, 6}

	// ind, err := searchInsertPosition(numbers, 7)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "error: %v\n", err)
	// }
	// fmt.Println(ind)
}

func searchInsertPosition(numbers []int, number int) (int, error) {
	low := 0
	high := len(numbers) - 1

	for {

		mid := low + (high-low)/2

		if low > high {
			return mid, nil
		}

		guess := numbers[mid]

		if guess == number {
			return mid, nil
		} else if guess > number {
			high = mid - 1
		} else {
			low = mid + 1
		}

	}
}

func binarySearch(nums []int, target int) (int, error) {
	low := 0
	high := len(nums) - 1

	for low <= high {
		mid := low + (high-low)/2
		guess := nums[mid]

		if guess == target {
			return mid, nil
		}

		if guess > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1, fmt.Errorf("number %d not found", target)
}

func firstEntry(nums []int, target int) int {
	low := 0
	high := len(nums) - 1
	result := -1

	for low <= high {
		mid := low + (high-low)/2

		if nums[mid] == target {
			result = mid
			high = mid - 1
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return result
}
