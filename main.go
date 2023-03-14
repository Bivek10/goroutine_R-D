package main

import (
	"sync"

	"example.com/go_routine_example/helpers"
)

var mu = &sync.Mutex{}
var wg = &sync.WaitGroup{}

func main() {

	//runtime.GOMAXPROCS(500)
	//fmt.Println("thread running", runtime.GOMAXPROCS(-1))
	//helpers.AsyncGreet(wg)
	//helpers.AsyncGetStatusCode(wg)
	//helpers.LaunchFunc()
	//helpers.ArraySum(wg, mu)
	//helpers.ArraySumFunc1(wg, mu)
	helpers.BasicChannel(wg)

}
