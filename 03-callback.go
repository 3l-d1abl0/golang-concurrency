package main

import (
	"fmt"
	"strings"
	"sync"
)

var wg sync.WaitGroup

func toUpperSync(word string, f func(string)) {

	f(strings.ToUpper(word))
}

func toLowerAsync(word string, f func(string)) {

	go func() {
		f(strings.ToLower(word))
	}()
}

func main() {
	//Synchronous Callback
	toUpperSync("Hello Callbacks !", func(v string) {
		fmt.Printf("Callback: %s\n", v)
	})

	//Async Callback
	wg.Add(1)
	toLowerAsync("Async Callback !", func(v string) {
		fmt.Printf("Callback : %s\n", v)
		wg.Done()
	})
	fmt.Println("Waiting ...:")
	wg.Wait()
}
