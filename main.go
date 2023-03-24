package main

import (
	"flag"
	"log"
	"net/http"
	"sync"

	"example.com/go_routine_example/helpers"
)

var addr = flag.String("addr", ":8080", "http service address")

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "./frontend/index.html")
}

func main() {

	// flag.Parse()
	// hub := helpers.NewHub()
	// go hub.Run()
	// http.HandleFunc("/", serveHome)
	// http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
	// 	helpers.ServeWs(hub, w, r)
	// })
	// err := http.ListenAndServe(*addr, nil)
	// if err != nil {
	// 	log.Fatal("ListenAndServe: ", err)
	// }

	ws := &sync.WaitGroup{}
	ws.Add(2)
	helpers.WriteMessage(ws)
	helpers.ReadMessage(ws)
	ws.Wait()
}
