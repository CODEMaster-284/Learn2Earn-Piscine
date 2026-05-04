package piscine

import "sort"

func Abort(a, b, c, d, e int) int {
	nums := []int{a, b, c, d, e}

	sort.Ints(nums)

	return nums[2]
}
