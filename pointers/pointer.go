package main

import "fmt"

func main() {
	age := 32 // regular variable

	// var agePointer *int // decalring a pointer
	agePointer := &age // assigning address of a variable

	fmt.Println("Age: ", age)
	fmt.Println("agePointer: ", agePointer)
	fmt.Println("agePointerValue: ", *agePointer) // getting value from pointer

	fmt.Println("Getting adult years: ", getAdultYears(agePointer)) // passing pointer to a function so that it does not create a copy of that variable
	mutateAge(agePointer)
	fmt.Println("getting mutated age value: ", age)
}
func getAdultYears(agePointer *int) int {
	return *agePointer - 18
}

func mutateAge(agePointer *int) {
	*agePointer = *agePointer - 18 // mutating age variable directly
}
