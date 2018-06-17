package main

import (
	"github.com/strider2038/go-examples/di"
)

func main() {
	runExampleDI()
}

func runExampleDI() {
	var serviceFactory di.ServiceFactory
	serviceFactory = di.ConcreteServiceFactory{}
	client := di.NewClient(serviceFactory)
	client.Run()
}
