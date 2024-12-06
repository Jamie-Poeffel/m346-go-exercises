package main

import "fmt"

func main() {
	// TODO: declare a type for Student (with first and last name)
	type Student struct {
		firstname string
		lastName  string
	}
	// TODO: declare a type for Class (consisting of multiple students)

	type Class map[int]Student

	ClassA := Class{
		1: Student{firstname: "Kush", lastName: "Srigiri"},
		2: Student{firstname: "Pascal", lastName: "Buehler"},
		3: Student{firstname: "Mike", lastName: "Bachman"},
	}

	ClassB := Class{
		1: Student{firstname: "Luis", lastName: "Allamand"},
		2: Student{firstname: "Julian", lastName: "Amschwanden"},
		3: Student{firstname: "John", lastName: "Doe"},
	}
	// TODO: declare a map of modules being attended by multiple classes
	type Modules map[int]Class

	// Initialisierung der Module
	modules := Modules{
		346: ClassA,
		117: ClassB,
		320: ClassA,
	}

	fmt.Println(modules)

	// TODO: output everything using fmt.Println()
}
