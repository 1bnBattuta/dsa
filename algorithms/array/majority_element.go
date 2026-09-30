package array

// A majority element is an element that appears in the array more
// than [n/2] times with n = len(arr)
// The following function uses Moore's voting algorithm
func Majority_element[T comparable](arr []T) (T, bool) {
	n := len(arr)

	if n == 0 {
		return *new(T), false
	}

	candidate := arr[0]
	count := 1

	for i := 2; i < n; i++ {
		if count == 0 {
			candidate = arr[i]
			count++
			continue
		}

		if arr[i] == candidate {
			count++
		} else {
			count--
		}
	}

	count = 0
	for i := range n {
		if candidate == arr[i] {
			count++
		}
	}

	if count >= n/2 {
		return candidate, true
	}
	return *new(T), false
}
