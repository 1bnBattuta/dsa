package main

import (
	"fmt"

	array "github.com/1bnBattuta/dsa/algorithms/array"
	math "github.com/1bnBattuta/dsa/algorithms/math"
)

func main() {
	math.Check_palindrome(123454321)

	arr := [10]int{2, 4, 5, 20, 9, 11, 4, 19, 88, 1}
	array.Reverse_array_in_groups(arr[:], 7)
	fmt.Println(arr)

}
