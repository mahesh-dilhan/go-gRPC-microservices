package main

import "net/rpc"

func main() {
	client, _ := rpc.DialHTTP("tcp", "127.0.0.1:9999")

}
