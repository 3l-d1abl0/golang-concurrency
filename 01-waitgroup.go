/*
wait groups synchronises the goroutines and wait for each one to finish
*/
package main

import "sync"

func main() {

	var wg sync.WaitGroup

	//waits for 1 goroutine
	wg.Add(1)

	go func() {
		println("Hello !")
		//wg.Done() //Time to exit
		wg.Add(-1)
	}()

	wg.Wait()

}
