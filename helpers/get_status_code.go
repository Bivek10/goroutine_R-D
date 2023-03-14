package helpers

import (
	"fmt"
	"net/http"
	"sync"
)

func GetStatusCode(endpoint string, wg *sync.WaitGroup) {
	defer wg.Done()
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("Something went wrong!")
	}
	fmt.Printf("%s  status code is: %d\n", endpoint, res.StatusCode)
}

func AsyncGetStatusCode(wg *sync.WaitGroup) {
	websiteList := []string{
		"https://google.com",
		"https://go.dev",
		"https://lco.dev",
		"https://github.com",
		"https://fb.com",
	}

	for _, v := range websiteList {
		wg.Add(1)
		go GetStatusCode(v, wg)
		wg.Wait()
	}
}
