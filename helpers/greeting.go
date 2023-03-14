package helpers

import (
	"fmt"
	"sync"
)

func greeting(greet string, ws *sync.WaitGroup) {
	defer ws.Done()
	for i := 0; i < 5; i++ {
		fmt.Println(greet)
	}
}

func AsyncGreet(ws *sync.WaitGroup) {
	ws.Add(1)
	go greeting("Hello", ws)
	ws.Wait()
}


