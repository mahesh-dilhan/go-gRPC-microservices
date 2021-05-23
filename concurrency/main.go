package main

import (
	"fmt"
	"net/http"
	"os"
	"sync"
)

var wg sync.WaitGroup
var mut sync.Mutex

func invokeExternalAPI(url string) {
	defer wg.Done()
	s, e := http.Get(url)
	if e != nil {
		fmt.Println("Unable to reach")
	}
	mut.Lock()
	fmt.Printf("[%d] %s\n", s.StatusCode, url)
	mut.Unlock()
}

func main() {

	if len(os.Args) < 2 {
		panic("Unable to invoke")
	}

	for _, url := range os.Args[1:] {
		go invokeExternalAPI("https://" + url)
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println("Concurrency ")

}
