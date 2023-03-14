package helpers

import (
	"fmt"
	"sync"
)

func BasicChannel(wg *sync.WaitGroup) {
	channel := make(chan int)
	wg.Add(2)

	//receiver goutine function
	go func(ch <-chan int) {
		defer wg.Done()

		for i := range ch {
			fmt.Println("received value: ", i)
		}

	}(channel)

	//sender goroutine function
	go func(ch chan<- int) {
		defer wg.Done()
		//send 42 to channel
		for j := 0; j <= 50; j++ {
			ch <- j
		}
		close(ch)

	}(channel)

	wg.Wait()
}
