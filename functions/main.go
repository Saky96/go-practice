package main

import "fmt"

func main() {
	//passingFuncAsParamAndValues()
	//closuresCalling()
	//recursionCalling()
	variadicFuncCalling()
}

func passingFuncAsParamAndValues() {
	numbers := []int{4, 5, 6}
	numbers2 := []int{9, 8}

	fmt.Println(transform(&numbers, double)) // passing function as a parameter eg. double
	fmt.Println(transform(&numbers, triple))

	fmt.Println(transform(&numbers, validate(numbers)))
	fmt.Println(transform(&numbers2, validate(numbers2))) //Validates which function to call eg: triple here

	fmt.Println(transform(&numbers, func(number int) int { //Syntax of passing an anonymous function as parameter instead of a named function
		return number * 2
	}))
}

func closuresCalling() {
	numbers := []int{4, 5, 6}
	doubleTransformer := factoryOfTransformers(2)
	TripleTransformer := factoryOfTransformers(3)

	fmt.Println(transform(&numbers, doubleTransformer))
	fmt.Println(transform(&numbers, TripleTransformer))
}

func recursionCalling() {
	fmt.Println(factorialLoop(5))
	fmt.Println(factorialRecursion(5))
}

func variadicFuncCalling() {
	fmt.Println(sumUp("grocery sum", 1, 2, 3, 4, 5)) // here we passed any number of parameters

	artItems := []int{6, 7, 8, 9, 10}
	fmt.Println(sumUp("art Items sum", artItems...)) // here we passed a slice that is converted into the parameters using ...
}
