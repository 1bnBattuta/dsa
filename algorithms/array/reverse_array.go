package array

func Revere_array[T any](arr []T) {
	n := len(arr)
	var temp T

	for i := 0; i < n/2; i++ {
		temp = arr[i]
		arr[i] = arr[n-1-i]
		arr[n-1-i] = temp
	}
}
