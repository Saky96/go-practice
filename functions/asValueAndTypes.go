package main

type transformFunc func(int) int //type declaration for func type

func transform(numbers *[]int, transformFn transformFunc) []int { // transformFn is a func passed as an argument

	transformedSlice := make([]int, len(*numbers))

	for index, number := range *numbers {
		transformedSlice[index] = transformFn(number)
	}

	return transformedSlice
}

func double(num int) int {
	return num * 2
}
func triple(num int) int {
	return num * 3
}

func validate(numbers []int) transformFunc {
	if len(numbers) < 2 {
		return double
	}
	return triple
}
