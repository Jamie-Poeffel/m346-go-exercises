package main

import (
	"fmt"
	"math"
)
func computeDiscriminant(a float64, b float64, c float64) float64 {
	return (math.Pow(b, 2) - 4 * a * c)
}
// TODO: implement the function computeQuadraticFormula
func computeQuadraticFormula(a float64, b float64, c float64){
	D := computeDiscriminant(a, b, c)

	switch {
		case D > 0:
			sqrtD := math.Sqrt(D)
			x1 := (-b + sqrtD) / (2 * a)
			x2 := (-b - sqrtD) / (2 * a)
			fmt.Printf("x1 = %.2f, x2 = %.2f\n", x1, x2)
		case D == 0:
			x := -b / (2 * a)
			fmt.Printf("x = %.2f\n", x)
		default: 
			fmt.Println("keine Lösung")
	}
	
	
}

func main() {
	// TODO: call the function computeQuadraticFormula
	computeQuadraticFormula(10, 20, 30) // keine Lösung
	computeQuadraticFormula(24, 14, 21) // keine Lösung
	computeQuadraticFormula(1 ,61, 34) 	// x1 = -0.56, x2 = -60.44
	computeQuadraticFormula(13, 25, 76) // keine Lösung

	computeQuadraticFormula(3, 4 ,1)	// x1 = -0.33, x2 = -1.00
	computeQuadraticFormula(2, 4, 2)	// x = -1.00
	computeQuadraticFormula(3, 4, 2)	// keine Lösung
}
