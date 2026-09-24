package main

import (
	"fmt"

	math "github.com/1bnBattuta/dsa/algorithms/math"
)

func main() {
	fmt.Println(math.Check_prime_fermat(1000000009, 5))
	fmt.Println(math.Check_prime_fermat(1000000008, 5))
}
