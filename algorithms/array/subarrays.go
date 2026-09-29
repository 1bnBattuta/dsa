package array

func Subarrays[T any](arr []T) [][]T {
	n := len(arr)
	result := make([][]T, 0, n*(n+1)/2)

	for i := 0; i < n; i++ {
		for j := i + 1; j <= n; j++ {
			result = append(result, arr[i:j])
		}
	}
	return result
}
