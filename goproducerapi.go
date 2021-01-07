package main

import ("net/http"
        "encoding/json"
)

func main() {
 http.HandleFunc("/greet/",GreetHandler)
 http.ListenAndServe(":8080", nil)
 
}

type Message struct {
 Response string
}

func GreetHandler(res http.ResponseWriter, req *http.Request) {
  
  m := Message{Response : "Server Response, build 2021-01-06"}
  b, err := json.Marshal(m)

  if err != nil {
   panic(err)
 }
  
 res.Write(b)
}

