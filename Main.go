package main

import (
	"fmt"
	"math"
)

func main() {
	var R int = 65
	var P float64 = 3.14
	var L float64 = (2 * P * float64(R))
	var S float64 = (P * (math.Pow(float64(R), 2)))
	fmt.Println(L, S)
}
