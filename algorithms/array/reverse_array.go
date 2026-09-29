package array

func Reverse_array[T any](arr []T) {
	n := len(arr)
	var temp T

	for i := 0; i < n/2; i++ {
		temp = arr[i]
		arr[i] = arr[n-1-i]
		arr[n-1-i] = temp
	}
}

func Reverse_array_in_groups[T any](arr []T, k int) {
	len := len(arr)
	n := len / k

	for i := 0; i < n; i++ {
		Reverse_array(arr[(i * k):((i + 1) * k)])
	}

	Reverse_array(arr[n*k : len])
}
