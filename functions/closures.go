package main

// here closures parameter factor is accessible by the anonymous function it is returning
func factoryOfTransformers(factor int) func(int) int {
	return func(number int) int { // returning a function as a value
		return factor * number // factor is enclosed inside the function and is accessible by this anonymous function
	}
}
