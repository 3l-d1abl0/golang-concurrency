package main

import "fmt"

/*
Channels Blocks:

If a channel is being listned to,
It expects to get a value.
I will continue to listen until it gets a value.

If a channel is of Size 1 and has been written once.

It will block if we try to write one more value,
Until someone reads the existing value
*/
func main() {

	//creating bidirectional Channel
	channel := make(chan string)

	go func() {
		//channel <- "Hello World !" //sending message in channel
		//blocking execution of go routine until message is received
		println("Read !")
	}()

	println("Reading Message !")
	message := <-channel //extracting message off channel
	fmt.Println(message)
	fmt.Println("--exiting--")
}
