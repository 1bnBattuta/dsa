package math

func Power(b float64, e int) float64 {
	if e == 0 || b == 1.0 {
		return 1
	}

	if e < 0 {
		return 1 / Power(b, -e)
	}

	if e%2 == 0 {
		return Power(b, e/2) * Power(b, e/2)
	} else {
		return b * Power(b, (e-1)/2) * Power(b, (e-1)/2)
	}
}
