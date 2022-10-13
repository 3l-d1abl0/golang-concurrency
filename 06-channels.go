package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	channel := make(chan string)
	wg.Add(1)

	go func() {
		channel <- "Hello"
		println("Finishing goroutine ")
		wg.Done()
	}()

	time.Sleep(time.Second)

	message := <-channel
	fmt.Println(message)
	wg.Wait()

}

//can have passive emitters and passive listeners
