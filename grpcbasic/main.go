package main

import (
	"fmt"
	"github.com/mahesh-dilhan/grpcbasic/who"
)

func main() {
	w := who.NewWHO()
	fmt.Println(w)
}
