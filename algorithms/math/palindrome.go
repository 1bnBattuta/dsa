package math

func Check_palindrome(n int) bool {
	reverse := 0

	if n < 0 {
		n = -n
	}

	temp := n

	for temp != 0 {
		reverse = (reverse * 10) + (temp % 10)
		temp = temp / 10
	}

	return reverse == n
}
