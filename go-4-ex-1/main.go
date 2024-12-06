package main

import (
	"fmt"
	"os"
)

// TODO: implement the function computeGrade
func computeGrade(points float32, maxpoints float32) {
	if (points < 0 || maxpoints < 0) {
		fmt.Fprintln(os.Stderr, "Points have to be positive")
		return
	}
	if (points > maxpoints) {
		fmt.Fprintln(os.Stderr, "Points cannot be bigger than max points")
		return
	}

	note := (points / maxpoints) * 5 + 1
	fmt.Println(note)
}

func main() {
	// TODO: call the function computeGrade
	computeGrade(17.5, 28.0)
	computeGrade(20, 10)
	computeGrade(-3, 20)
}
