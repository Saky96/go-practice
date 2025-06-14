package main

import "fmt"

// It is a variadic function that can take any number of arguments and convert them into a slice to be used
func sumUp(listName string, numbers ...int) int {
	fmt.Println("It is a list of: " + listName)
	sum := 0
	for _, val := range numbers {
		sum += val
	}

	return sum
}
