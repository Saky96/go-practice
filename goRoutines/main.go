package main

import (
	"github.com/Saky96/go-practice/goRoutines/examples"
)

func main() {

	//linearExecutions()

	//parallelExecutions()
	parallelExecutions2()

}

func linearExecutions() {
	examples.SimpleGreet("Hello Beta!")
	examples.SimpleGreet("Kya Haal")
	doneChan := make(chan bool)
	go examples.LongGreet("ab kya karu ", doneChan, 3)
	examples.SimpleGreet("Koi na")
	<-doneChan // Wait for the goroutine to finish
}

func parallelExecutions() {
	doneChan := make(chan bool)
	go examples.LongGreet("Hello Beta!", doneChan, 4)
	go examples.LongGreet("Kya Haal", doneChan, 1)
	go examples.LongGreet("ab kya karu ", doneChan, 2)
	go examples.LongGreet("Koi na", doneChan, 8)
	<-doneChan // Wait for the goroutine to finish
	<-doneChan // Wait for the goroutine to finish
	<-doneChan // Wait for the goroutine to finish
	<-doneChan // Wait for the goroutine to finish
}

//----------------------------Altenative Appoach-----------------------------

func parallelExecutions2() {
	// dones := make([]chan bool, 4)
	done := make(chan bool)

	// dones[0] = make(chan bool)
	go examples.LongGreet("Nice to meet you!", done, 2)
	// dones[1] = make(chan bool)
	go examples.LongGreet("How are you?", done, 1)
	// dones[2] = make(chan bool)
	go examples.LongGreet("How ... are ... you ...?", done, 5)
	// dones[3] = make(chan bool)
	go examples.LongGreet("I hope you're liking the course!", done, 8)

	// for _, done := range dones {
	// 	<-done
	// }

	for range done {
		// fmt.Println(doneChan)
	}
}
