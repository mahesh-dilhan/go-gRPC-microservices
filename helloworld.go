package main

import ("fmt"
        "net/http"
)

func main() {
 http.HandleFunc("/", HelloServer)
 http.ListenAndServe(":80", nil)

}

func HelloServer(res http.ResponseWriter,  req *http.Request) {
   fmt.Fprintf(res, "Hello, %s!", req.URL.Path[1:])
   fmt.Println("")

}
