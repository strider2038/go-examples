package main

import (
	"flag"
	"fmt"
	"github.com/strider2038/go-examples/di"
	"github.com/strider2038/go-examples/jsonServer"
	"github.com/strider2038/go-examples/jsonrpc"
	"github.com/strider2038/go-examples/jsonrpc2"
)

func main() {
	var exampleId string
	flag.StringVar(&exampleId, "example", "di", "id of the running example")
	flag.Parse()
	fmt.Println("Running example:", exampleId)

	switch exampleId {
	case "di":
		runExampleDI()
	case "jsonrpc":
		jsonrpc.RunJSONRPCServer()
	case "jsonrpc2":
		jsonrpc2.RunJSONRPC2Server()
	case "json":
		jsonServer.RunJsonServer()
	}
}

func runExampleDI() {
	var serviceFactory di.ServiceFactory
	serviceFactory = di.ConcreteServiceFactory{}
	client := di.NewClient(serviceFactory)
	client.Run()
}
