package main

import "fmt"

const (
	Lower = 1
	Upper = 30
)

func main() {
	// TODO: Implement FizzBuzz using a for loop from Lower to Upper.

	for i := Lower; i <= Upper; i++ {
		fmt.Print(i, ". ")
		switch {
			case i%5 == 0 && i % 3 == 0: {
				fmt.Println("FizzBuzz")
			}
			case i % 3 == 0: {
				fmt.Println("Fizz")
			}
			case i % 5 == 0: {
				fmt.Println("Buzz")
			}
			default: {
				fmt.Println(i)
			}
		}
		// if i%5 == 0 && i % 3 == 0 {
		// 	fmt.Println("FizzBuzz")
		// }else if i % 3 == 0 {
		// 	fmt.Println("Fizz")
		// }else if i % 5 == 0{
		// 	fmt.Println("Buzz")
		// }else {
		// 	fmt.Println(i)
		// }
	}
}
