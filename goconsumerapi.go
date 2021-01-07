package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
)

func main() {
    res, err := http.Get("http://localhost:8080/greet/")

    if err !=nil {
	panic(err)
    }

   body, err := ioutil.ReadAll(res.Body)    

    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(string(body))

}
