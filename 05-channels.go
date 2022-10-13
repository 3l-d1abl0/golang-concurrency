package main

import "fmt"

func main() {

	//creating bidirectional Channel
	channel := make(chan string)

	go func() {
		channel <- "Hello World !" //sending message in channel
		//blocking execution of go routine until message is received
		println("Read !")
	}()

	println("Reading Message !")
	message := <-channel //extracting message off channel
	fmt.Println(message)
	fmt.Println("--exiting--")
}
