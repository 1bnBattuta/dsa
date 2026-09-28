package math

func order(n int) int {
	o := 0
	for n != 0 {
		o++
		n = n / 10
	}
	return o
}

func power(x int, y int) int {
	if y == 0 {
		return 1
	}
	if y%2 == 0 {
		return power(x, y/2) * power(x, y/2)
	}
	return x * power(x, y/2) * power(x, y/2)
}

func Check_armstrong(n int) bool {

	sum := 0
	order := order(n)
	temp := n

	for temp > 0 {
		r := temp % 10
		sum += power(r, order)
		temp = temp / 10
	}

	return (sum == n)
}
