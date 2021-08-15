package main

import (
	"encoding/json"
	"fmt"
	"github.com/mahesh-dilhan/grpcbasic/proto/who/region"
)

func main() {

	cntry := &region.Region{
		K: &region.Key{
			Name:  "Asia/Pasific",
			State: "NY",
		},
		PositiveCases: 1000000,
	}

	//fmt.Println(cntry)

	data, err := json.Marshal(cntry)
	if err != nil {
		fmt.Println("Error in marshell", err)
	}
	fmt.Println(data)

}
