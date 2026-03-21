package main

func twoSum(nums []int, target int) []int {

	m := make(map[int]int)

	for i, x := range nums {
		y := target - x

		val, ok := m[y]
		if ok {
			return []int{val, i}
		} else {
			m[x] = i
		}
	}

	return []int{}
}
