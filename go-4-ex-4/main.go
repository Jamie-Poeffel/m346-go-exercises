package main

import (
	"fmt"
)

type Celsius float64 

func (C Celsius) ConvertToFahrenheit() Fahrenheit{
	return convertCelsiusToFahrenheit(C)
}

type Fahrenheit float64

func (F Fahrenheit) ConvertToCelsius() Celsius{
	return convertFahrenheitToCelsius(F)
}


// TODO: implement the function convertCelsiusToFahrenheit
func convertCelsiusToFahrenheit(C Celsius) Fahrenheit {
	//F = C * (9/5) + 32
	return Fahrenheit(C*(9.0/5.0) + 32.0)
}
// TODO: implement the function convertFahrenheitToCelsius
func convertFahrenheitToCelsius(F Fahrenheit) Celsius {
	// C = (F - 32) * (5/9)
	return Celsius((F - 32.0) * (5.0/9.0))
}

func main() {
	// TODO: call the function convertCelsiusToFahrenheit
 
	fmt.Println(convertCelsiusToFahrenheit(17.0))	//62.6
	fmt.Println(convertCelsiusToFahrenheit(5.7))	//42.26
	fmt.Println(convertCelsiusToFahrenheit(6.9))	//44.42

	fmt.Println()
	// TODO: call the function convertFahrenheitToCelsius

	fmt.Println(convertFahrenheitToCelsius(62.6))	//17
	fmt.Println(convertFahrenheitToCelsius(42.26))	//5.69999999999
	fmt.Println(convertFahrenheitToCelsius(44.42))	//6.90000000001

	fmt.Println()

	var cozy Celsius = 23.0
	fmt.Println(cozy.ConvertToFahrenheit()) //73.4

	var cold Fahrenheit = 15.3
	fmt.Println(cold.ConvertToCelsius()) //-9.2777777779	
}
