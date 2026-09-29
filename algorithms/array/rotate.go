package array

func reverse_array[T any](arr []T) {
	n := len(arr)
	var temp T

	for i := 0; i < n/2; i++ {
		temp = arr[i]
		arr[i] = arr[n-1-i]
		arr[n-1-i] = temp
	}
}

func Rotate_array[T any](arr []T, d int) {
	n := len(arr)
	d %= n

	if d == 0 {
		return
	}

	reverse_array(arr[:])
	reverse_array(arr[:d])
	reverse_array(arr[d:])
}
