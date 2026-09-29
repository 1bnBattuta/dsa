package main

import (
	"fmt"

	array "github.com/1bnBattuta/dsa/algorithms/array"
	math "github.com/1bnBattuta/dsa/algorithms/math"
)

func main() {
	math.Check_palindrome(123454321)

	a := [4]int{1, 2, 3, 4}
	fmt.Println(array.Subsequences(a[:]))
	fmt.Println(array.Subarrays(a[:]))

}
