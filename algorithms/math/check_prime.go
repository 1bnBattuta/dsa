package math

import (
	"math/rand"
)

func gcd(a int, b int) int {
	if a < b {
		return gcd(b, a)
	} else if a%b == 0 {
		return b
	} else {
		return gcd(b, a%b)
	}
}

func power_remainder(a int, n int, p int) int {
	res := 1
	a = a % p

	for n > 0 {
		if n%2 == 1 {
			res = (res * a) % p
		}
		n = n / 2
		a = (a * a) % p
	}
	return res
}

func Check_prime_square_root(n int) bool {

	if n == 1 {
		return false
	}
	if n < 4 {
		return true
	}

	if n%2 == 0 || n%3 == 0 {
		return false
	}

	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}

	return true
}

// The following is a probabilistic method based on Fermat's little theorem.
// If n is prime, then for all i in [2, n-2] we have: (*) i^(n-1) % n = 1
// If n is composite, the probability that all checked numbers satisfy (*)
// is low and can be further reduced by increasing the number of iterations
// k.
// This algo fails for carmichael numbers.
func Check_prime_fermat(n int, k int) bool {

	if n <= 1 || n == 4 {
		return false
	}
	if n <= 3 {
		return true
	}

	var x int
	for range k {
		x = rand.Intn(n-3) + 2

		if gcd(n, x) != 1 {
			return false
		}

		if power_remainder(x, n-1, n) != 1 {
			return false
		}
	}
	return true
}
