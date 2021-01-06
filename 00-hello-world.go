package main

import "time"

func main() {

	go helloWorld()

	//Anon function
	go func(msg string) {
		println(msg)
	}("Hello World from Anon !")

	messagePrint := func(msg string) {
		println(msg)
	}

	go messagePrint("Hello World from messagePrint!")
	go messagePrint("Hello from Goroutine !")

	time.Sleep(time.Second)
	println("--exiting--")
}

func helloWorld() {

	println("Hello World")
}
