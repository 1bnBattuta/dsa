package array

func Subsequences[T any](arr []T) [][]T {
	if len(arr) == 0 {
		return nil
	}

	elem := arr[0]
	rest := Subsequences(arr[1:])

	result := make([][]T, 0, len(rest)*2)

	for _, sub := range rest {
		// subsequences without elem
		result = append(result, sub)
		// subsequence with elem prepended
		withElem := make([]T, 0, len(sub)+1)
		withElem = append(withElem, elem)
		withElem = append(withElem, sub...)
		result = append(result, withElem)
	}

	result = append(result, []T{elem})
	return result
}
