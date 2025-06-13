package main

import (
	"example/arrays/exercises"
	"fmt"
)

func main() {

	//arrayBasics()
	exercises.Exe1()
}

func arrayBasics() {
	listOfInts := [4]float64{9.08, 5.65, 3.76, 0.892}          // float arrays
	productNames := []string{"Appa", "Backr", "Cheta", "tree"} // dynamic String array
	subProducts := productNames[1:3]                           // Slicing and array and selecting from index 1 till index 3 (excluded)

	fmt.Println("Hello Arrays!")
	fmt.Println(listOfInts)
	fmt.Println(productNames)
	fmt.Println(subProducts)
	fmt.Println(productNames[1:])                   //Starting from index 1 till last element
	fmt.Println(productNames[:2])                   //Starting from beginning going till the 2nd index
	fmt.Println(len(subProducts), cap(subProducts)) //len() gives length of current slice, capacity gives capacity according to the array from which slice is made up of

	productNames = append(productNames, "Jabba")
	fmt.Println(productNames)
}
