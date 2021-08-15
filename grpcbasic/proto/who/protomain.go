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
		Status:        region.Region_MILD,
	}

	//fmt.Println(cntry)

	data, err := json.Marshal(cntry)
	if err != nil {
		fmt.Println("Error in marshell", err)
	}
	fmt.Println(data)

	r := &region.Region{}

	err = json.Unmarshal(data, r)
	if err != nil {
		fmt.Println("Error unmarshal", err)
	}
	fmt.Println(r)
	fmt.Println(r.GetStatus())

}
