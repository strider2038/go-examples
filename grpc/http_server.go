package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/strider2038/go-examples/grpc/messaging"
	"google.golang.org/grpc"
)

func RunHTTPServer(grpcAddress string, proxyAddress string) {
	grpcConnection, err := grpc.Dial(grpcAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalln("failed to connect to grpc", err)
	}
	defer grpcConnection.Close()

	grpcGWMux := runtime.NewServeMux()

	err = messaging.RegisterMessengerHandler(
		context.Background(),
		grpcGWMux,
		grpcConnection,
	)
	if err != nil {
		log.Fatalln("failed to start HTTP server", err)
	}

	mux := http.NewServeMux()
	// отправляем в прокси только то что нужно
	mux.Handle("/", grpcGWMux)

	fmt.Println("starting HTTP server at " + proxyAddress)
	log.Fatal(http.ListenAndServe(proxyAddress, mux))
}
