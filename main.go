package main

import (
	"net/http"
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
	// helpers.BasicChannel(wg)
	//http.HandleFunc("/ws", helpers.Handler)
	setupAPI()

}

func setupAPI() {
	http.Handle("/", http.FileServer(http.Dir("./frontend")))
	http.HandleFunc("/ws", helpers.NewManager().ServeWS)

	http.ListenAndServe(":8080", nil)
}
