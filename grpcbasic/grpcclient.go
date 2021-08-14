package main

import (
	"fmt"
	"github.com/mahesh-dilhan/grpcbasic/who"
	"net/rpc"
)

func main() {
	client, _ := rpc.DialHTTP("tcp", "127.0.0.1:9999")
	var covidcountry who.Country

	id := new(who.Key{
		Name:  "USA",
		State: "NY",
	})
	payload := new(who.Country{
		K:             id,
		PositiveCases: 100,
	})

	if err := client.Call("WHO.Add", payload, &covidcountry); err != nil {
		fmt.Println("unable to Add", err)
	} else {
		fmt.Printf("Sucess '%v'\n", covidcountry.K)
	}

}
