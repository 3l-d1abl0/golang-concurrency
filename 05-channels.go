package main

import "fmt"

func main() {

	channel := make(chan string)

	go func() {
		channel <- "Hello World !" //sending message in channel
	}()

	message := <-channel //extracting message off channel
	fmt.Println(message)
	fmt.Println("--exiting--")
}
