package main

import (
	"fmt"
	"github.com/mahesh-dilhan/grpcbasic/who"
	"io"
	"net/http"
	"net/rpc"
)

func main() {
	w := who.NewWHO()
	fmt.Println(w)
	rpc.Register(w)

	rpc.HandleHTTP()
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		io.WriteString(res, "server alive...")
	})
	http.ListenAndServe(":9999", nil)
}
