// http://www.nada.kth.se/~snilsson/concurrency/
package main

import (
	"fmt"
	"time"
)

func main() {
	Publish("A goroutine starts a new thread of execution.", 5*time.Second)
	fmt.Println("Let us hope the news will published before I leave.")

	// Wait for the news to be published.
	time.Sleep(10 * time.Second)

	fmt.Println("Ten seconds later: I am leaving now.")
}

// Publish prints text to stdout after the given time has expired.
// It does not block but returns right away.
func Publish(text string, delay time.Duration) {
	go func() {
		time.Sleep(delay)
		fmt.Println("BREAKING NEWS:", text)
	}() // Note the parentheses. We must call the anonymous function.
}
