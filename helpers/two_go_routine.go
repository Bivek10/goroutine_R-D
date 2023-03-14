/*
Create a program that launches two goroutines.
One of the goroutines should print out "hello" every 1 second,
and the other goroutine should print out "world"
every 2 seconds. The program should terminate after 10 seconds.
*/

package helpers

import (
	"fmt"
	"time"
)

func LaunchFunc() {
	fmt.Println("---------------Start Program--------------")

	go printHello()
	go printWorld()

	time.Sleep(10 * time.Second)
	fmt.Println("----------------End Program---------------")
}

func printHello() {
	for {
		fmt.Println("Hello")
		time.Sleep(1 * time.Second)
	}
}

func printWorld() {
	for {
		fmt.Println("World!")
		time.Sleep(2 * time.Second)
	}
}
