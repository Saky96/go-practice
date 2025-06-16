package examples

import (
	"fmt"
	"time"
)

func SimpleGreet(greet string) {
	fmt.Println(greet)
}

func LongGreet(greet string, doneChan chan bool, sleepTime time.Duration) {
	time.Sleep(sleepTime * time.Second)
	fmt.Println("Time gone " + greet + " ...")
	doneChan <- true
	close(doneChan)
}
