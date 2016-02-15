// http://www.nada.kth.se/~snilsson/concurrency/
package main

import "fmt"

type Sushi string

func main() {
	var ch <-chan Sushi = Producer()
	for s := range ch {
		fmt.Println("Consumed", s)
	}
}

func Producer() <-chan Sushi {
	ch := make(chan Sushi)
	go func() {
		ch <- Sushi("娴疯€佹彙銈�")  // Ebi nigiri
		ch <- Sushi("楫仺銈嶆彙銈�") // Toro nigiri
		close(ch)
	}()
	return ch
}
