package main

import (
	"fmt"
	"math"
)

type ShortSides struct {
	a float64
	b float64
}

// TODO: implement the function computeHypotenuse using math.Pow and math.Sqrt
func computeHypotenuse(a float64, b float64) {
	c := math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
	fmt.Println(c)
}

func Hypotenuse(s ShortSides) {
	c := math.Sqrt(math.Pow(s.a, 2) + math.Pow(s.b, 2))
	fmt.Println(c)
}


func main() {
	// TODO: call the function computeHypotenuse
	computeHypotenuse(19, 34)	// 38.948684188300895
	computeHypotenuse(22, 3)		// 22.20360331117452
	computeHypotenuse(34, 12)	// 36.05551275463989
	computeHypotenuse(15, 18)	// 23.430749027719962

	Hypotenuse(ShortSides{114, 32})	// 118.40608092492548
	Hypotenuse(ShortSides{2, 36})	// 36.05551275463989
	Hypotenuse(ShortSides{10, 12})	// 15.620499351813308
}
