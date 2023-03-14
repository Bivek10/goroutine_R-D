package main

import (
	"sync"

	"example.com/go_routine_example/helpers"
)

var mu = &sync.Mutex{}
var wg = &sync.WaitGroup{}

func main() {

	//helpers.AsyncGreet(wg)
	//helpers.AsyncGetStatusCode(wg)
	//helpers.LaunchFunc()
	//helpers.ArraySum(wg, mu)
	helpers.ArraySumFunc1(wg, mu)

}
