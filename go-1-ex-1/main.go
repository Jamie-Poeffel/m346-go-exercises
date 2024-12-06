package main

import "fmt"

func main() {
	// TODO: Declare and initialize the variables being used in the output!
	var firstName string = "Jamie"
	var lastName string = "Poeffel"

	dayOfBirth := 3
	monthOfBirth := 10
	yearOfBirth := 2007

	var numberOfSiblings = 1

	heightInMeters := 1.78

	var zodiacSign rune = '\u264e'

	fmt.Printf("Vor- und Nachname: %s %s\n", firstName, lastName)
	fmt.Printf("Geburtsdatum: %d.%d.%d\n", dayOfBirth, monthOfBirth, yearOfBirth)
	fmt.Printf("Anzahl Geschwister: %d\n", numberOfSiblings)
	fmt.Printf("Grösse (in Metern): %.2f\n", heightInMeters)
	fmt.Printf("Sternzeichen: %c\n", zodiacSign)
}
