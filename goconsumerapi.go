package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "encoding/json"
)

type Message struct {
 Response string
}
func main() {
    res, err := http.Get("http://localhost:8080/greet/")

    if err !=nil {
	panic(err)
    }

   body, err := ioutil.ReadAll(res.Body)    

    if err != nil {
        log.Fatal(err)
    }
   
    var msg Message
    json.Unmarshal(body, &msg)
    fmt.Println(msg)

}
