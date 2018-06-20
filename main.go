package main

import (
	"flag"
	"fmt"
	"github.com/strider2038/go-examples/di"
	"github.com/strider2038/go-examples/jsonrpc"
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
	}
}

func runExampleDI() {
	var serviceFactory di.ServiceFactory
	serviceFactory = di.ConcreteServiceFactory{}
	client := di.NewClient(serviceFactory)
	client.Run()
}
