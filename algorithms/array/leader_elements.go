package array

import (
	"golang.org/x/exp/constraints"
)

// A leader element is an element that is greater than or
// equal to all elements to its right
func Leader_elements[T constraints.Ordered](arr []T) []T {
	n := len(arr)
	res := make([]T, 0, n)

	if n == 0 {
		return res
	}

	res = append(res, arr[n-1])
	current_max := arr[n-1]

	for i := n - 2; i >= 0; i-- {
		if arr[i] >= current_max {
			res = append(res, arr[i])
			current_max = arr[i]
		}
	}

	return res
}
