package array

func Multiply_with_adjacent(arr []int) {
	prev := 1
	len := len(arr)

	for i := 0; i < len; i++ {
		curr := arr[i]
		next := 1
		if i < len-1 {
			next = arr[i+1]
		}

		arr[i] = prev * curr * next
		prev = curr
	}
}
