package main

import (
	"fmt"

	array "github.com/1bnBattuta/dsa/algorithms/array"
	math "github.com/1bnBattuta/dsa/algorithms/math"
)

func main() {
	math.Check_palindrome(123454321)

	a := [8]int{3, 1, 4, 6, 1, 1, 6, 1}
	fmt.Println(array.Majority_element(a[:]))
}
