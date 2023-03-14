package helpers

import (
	"fmt"
	"sync"
)

func BasicChannel(wg *sync.WaitGroup) {
	channel := make(chan int)
	wg.Add(2)
	//receiver goutine function
	go func() {
		defer wg.Done()
		//received value from channel

		i := <-channel
		fmt.Println("received value: ", i)
		//send new value back to channel
		channel <- 10

	}()

	//sender goroutine function
	go func() {
		defer wg.Done()
		//send 42 to channel
		channel <- 42
		fmt.Println("received value: ", <-channel)

	}()
	wg.Wait()
}
