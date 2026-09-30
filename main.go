package main

import (
	"fmt"

	array "github.com/1bnBattuta/dsa/algorithms/array"
	math "github.com/1bnBattuta/dsa/algorithms/math"
)

func main() {
	math.Check_palindrome(123454321)

	a := [8]int{-1, 2, 3, -4, 12, -7, 6, 10}
	array.Rearrange_by_sign(a[:])
	fmt.Println(a)
}
