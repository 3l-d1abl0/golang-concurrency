package main

import (
	"fmt"
	"strings"
	"sync"
)

var wg sync.WaitGroup

func toUpper(word string, f func(string)) {

	go func() {
		f(strings.ToUpper(word))
	}()
}

func toLower(word string, f func(string)) {

	go func() {
		f(strings.ToLower(word))
	}()
}

func main() {

	wg.Add(1)

	toUpper("1st Call", func(wrd string) {

		fmt.Printf("%s Calling toLower: \n", wrd)

		toLower("2nd Call", func(wrd string) {
			fmt.Printf("Callback: %s\n", wrd)
			wg.Done()
		})
	})

	println("Waiting ....")
	wg.Wait()
}
