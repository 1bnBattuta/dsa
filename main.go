package main

import (
	"fmt"

	array "github.com/1bnBattuta/dsa/algorithms/array"
	math "github.com/1bnBattuta/dsa/algorithms/math"
)

func main() {
	math.Check_palindrome(123454321)

	arr := [4]int{2, 4, 5, 20}
	array.Multiply_with_adjacent(arr[:])
	fmt.Println(arr)
	array.Revere_array(arr[:])
	fmt.Println(arr)

}
