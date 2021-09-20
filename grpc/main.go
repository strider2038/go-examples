package main

import "context"

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go RunServer(ctx, "127.0.0.1:8080")
	RunHTTPServer("127.0.0.1:8080", ":8081")
}
