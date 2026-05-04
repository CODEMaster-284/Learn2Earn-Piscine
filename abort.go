package piscine

func Abort(a, b, c, d, e int) int {
	arr := []int{a, b, c, d, e}
	n := len(arr)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if arr[i] > arr[j] {
				// Swap
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr[2]
}
