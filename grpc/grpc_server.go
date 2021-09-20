package main

import (
	"context"
	"fmt"
	"net"

	"github.com/strider2038/go-examples/grpc/messaging"
	"google.golang.org/grpc"
)

func RunServer(ctx context.Context, address string) error {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("failed to start net listener: %v", err)
	}

	go func(ctx context.Context, listener net.Listener) {
		server := newServer()
		go server.Serve(listener)

		select {
		case <-ctx.Done():
			server.GracefulStop()
		}
	}(ctx, listener)

	return nil
}

func newServer() *grpc.Server {
	server := grpc.NewServer()

	messaging.RegisterMessengerServer(server, NewMessenger())

	return server
}

