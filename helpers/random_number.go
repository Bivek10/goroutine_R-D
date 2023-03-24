// 4). Write a program that launches two goroutines.
// One of the goroutines should generate random numbers between 1 and 100 and send them to a channel.
//  The othesr goroutine should read numbers from the channel and print out whether the number is even or odd.

package helpers

import (
	"fmt"
	"math/rand"
	"sync"
)

func RandomNumber(wg *sync.WaitGroup) {
	cha := make(chan int)
	wg.Add(2)
	go func(ch <-chan int) {
		defer wg.Done()
		for {
			if val, ok := <-ch; ok {
				if val%2 == 0 {
					fmt.Println("received valued is even :", val)
				} else {
					fmt.Println("received valued is odd :", val)
				}
			} else {
				fmt.Println("Channel Closed")
				break
			}
		}

	}(cha)

	go func(ch chan<- int) {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			random := rand.Intn(100)
			ch <- random
		}
		close(ch)

	}(cha)

	wg.Wait()

}
