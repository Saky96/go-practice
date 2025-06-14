package main

func factorialLoop(num int) int {
	result := 1
	for i := num; i > 0; i-- {
		result = result * i
	}

	return result
}

func factorialRecursion(num int) int {
	if num == 1 {
		return 1
	}

	return num * factorialRecursion(num-1)
}
