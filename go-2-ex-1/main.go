package main

import "fmt"

type FullName struct {
	// TODO: add fields
	Vorname 	string
	Nachname 	string
}

// TODO: declare a structure for birth date
type birth struct {
	day 	int
	month	int
	year 	int
}

type Profile struct {
	// TODO: embed full name and birth date information
	NumberOfSiblings byte
	ZodiacSign       rune
	birth
	FullName
}

func main() {
	var me = Profile{
		// TODO: set name and birth date information
		NumberOfSiblings: 1,   // TODO: adjust
		ZodiacSign:       '\u264e', // TODO: adjust
		birth: birth{
			day: 3,
			month: 10,
			year: 2007,
		},
		FullName: FullName{
			Vorname: "Jamie",
			Nachname: "Poeffel",
		},
	}
	fmt.Println(me)

	fmt.Println("Siblings Before:", me.NumberOfSiblings)
	// TODO: imagine, you get a little brother or sister
	me.NumberOfSiblings++
	
	fmt.Println("Siblings After:", me.NumberOfSiblings)
}
