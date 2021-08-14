package main

import (
	"fmt"
	"github.com/mahesh-dilhan/grpcbasic/who"
	"net/rpc"
)

func main() {
	client, _ := rpc.DialHTTP("tcp", "127.0.0.1:9999")
	var covidcountry who.Country

	id := &who.Key{
		Name:  "USA",
		State: "NY",
	}
	payload := &who.Country{
		K:             *id,
		PositiveCases: 100,
	}

	payload2 := who.Country{
		K: who.Key{
			Name:  "SG",
			State: "SGK",
		},
		PositiveCases: 100,
	}

	if err := client.Call("WHO.Add", payload, &covidcountry); err != nil {
		fmt.Println("unable to Add", err)
	} else {
		fmt.Printf("Sucess '%v'\n", covidcountry.K)
	}

	if err := client.Call("WHO.Add", payload2, &covidcountry); err != nil {
		fmt.Println("unable to Add", err)
	} else {
		fmt.Printf("sucess '%v'\n", covidcountry.K)
	}
}
