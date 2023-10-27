package main

import (
	"time"
)

func main() {

	//Bidirectional Channel, used as bidirectional Channel
	helloCh := make(chan string, 1)
	//Bidirectional Channel, used as recieve from channel
	goodbyeCh := make(chan string, 1)
	//Bidirection Channel, used as send to channel
	quitCh := make(chan bool)

	//Setting up the recievers
	go receiver(helloCh, goodbyeCh, quitCh)

	go sendString(helloCh, "hello!")
	time.Sleep(time.Second)
	go sendString(goodbyeCh, "goodbye!")

	println("Waiting")
	println(<-quitCh)
	println("Exiting")

}

func sendString(ch chan<- string, s string) {
	//ch is send to channel
	ch <- s
}

func receiver(helloCh, goodbyeCh <-chan string, quitCh chan<- bool) {
	/*	helloCh - bidirectional Channel
		goodbyCh - receive from channel
		quitCh <-send to channel
	*/
	for {

		select {

		case msg := <-helloCh:
			println(msg)

		case msg := <-goodbyeCh:
			println(msg)

		case <-time.After(time.Second * 2):
			println("Nothing received in 2 seconds. Exiting")
			quitCh <- true
			break
		}
	}
}
