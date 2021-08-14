package main

import (
	"fmt"
	"github.com/mahesh-dilhan/grpcbasic/who"
	"net/rpc"
)

func main() {
	client, _ := rpc.DialHTTP("tcp", "127.0.0.1:9999")
	var covidcountry who.Country

	id := who.Key{
		Name:  "USA",
		State: "NY",
	}
	payload := who.Country{
		K:             id,
		PositiveCases: 100,
	}

	id2 := &who.Key{
		Name:  "SG",
		State: "SGK",
	}
	payload2 := &who.Country{
		K:             *id2,
		PositiveCases: 100,
	}

	id3 := new(who.Key)
	id3.Name = "IND"
	id3.State = "MUM"

	payload3 := who.Country{
		K:             *id3,
		PositiveCases: 1000,
	}

	//Async
	gCal := client.Go("WHO.Add", payload, &covidcountry, nil)
	replyCall, ok := <-gCal.Done

	if !ok {
		fmt.Println("unable to Add", replyCall)
	} else {
		fmt.Printf("Sucess '%v'\n", covidcountry.K)
	}

	////Add
	//if err := client.Call("WHO.Add", payload, &covidcountry); err != nil {
	//	fmt.Println("unable to Add", err)
	//} else {
	//	fmt.Printf("Sucess '%v'\n", covidcountry.K)
	//}

	if err := client.Call("WHO.Add", payload2, &covidcountry); err != nil {
		fmt.Println("unable to Add", err)
	} else {
		fmt.Printf("sucess '%v'\n", covidcountry.K)
	}

	if err := client.Call("WHO.Add", payload3, &covidcountry); err != nil {
		fmt.Println("unable to Add", err)
	} else {
		fmt.Printf("sucess '%v'\n", covidcountry.K)
	}

	//Get
	gcall := client.Go("WHO.Get", payload, &covidcountry, nil)
	replyCall, ok = <-gcall.Done
	if !ok {
		fmt.Println("unable to find ", replyCall)
	} else {
		fmt.Printf("found '%v' '%v'\n", covidcountry.K, covidcountry.PositiveCases)
	}

	//Get
	if err := client.Call("WHO.Get", payload, &covidcountry); err != nil {
		fmt.Println("unable to find ", err)
	} else {
		fmt.Printf("found '%v' '%v'\n", covidcountry.K, covidcountry.PositiveCases)
	}

	if err := client.Call("WHO.Get", payload2, &covidcountry); err != nil {
		fmt.Println("unable to find ", err)
	} else {
		fmt.Printf("found '%v' '%v'\n", covidcountry.K, covidcountry.PositiveCases)
	}

	if err := client.Call("WHO.Get", payload3, &covidcountry); err != nil {
		fmt.Println("unable to find ", err)
	} else {
		fmt.Printf("found '%v' '%v'\n", covidcountry.K, covidcountry.PositiveCases)
	}
}
