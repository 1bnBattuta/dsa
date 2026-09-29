package main

import (
	"fmt"

	array "github.com/1bnBattuta/dsa/algorithms/array"
	math "github.com/1bnBattuta/dsa/algorithms/math"
)

func main() {
	math.Check_palindrome(123454321)

	arr := [3]int{2, 4, 5}
	array.Multiply_with_adjacent(arr[:])
	fmt.Println(arr)
}
