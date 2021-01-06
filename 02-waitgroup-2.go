package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	goroutinesCount := 10

	wg.Add(goroutinesCount)

	for i := 0; i < goroutinesCount; i++ {
		go func(id int) {
			fmt.Printf("goroutine id: %d :: says Hello ! \n", id)
			wg.Done()
		}(i)
	}

	wg.Wait()
}
