package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"
)

var ErrConnectionRefused = errors.New("connection refused")

var start = time.Now()

type Service interface {
	Send(ctx context.Context, request string) (string, error)
	Connect(ctx context.Context) error
}

type ServiceImplementation struct {
	sendCount    int
	connectCount int
}

func (s *ServiceImplementation) Send(ctx context.Context, request string) (string, error) {
	s.sendCount++

	if s.sendCount > 1 {
		return "response", nil
	}

	return "", ErrConnectionRefused
}

func (s *ServiceImplementation) Connect(ctx context.Context) error {
	s.connectCount++

	if s.connectCount >= 5 {
		return nil
	}

	return ErrConnectionRefused
}

type ReconnectingService struct {
	wrapped Service

	maxAttempts int
	timeout     time.Duration
	multiplier  int
}

func (reconnector *ReconnectingService) Send(ctx context.Context, request string) (string, error) {
	for {
		response, err := reconnector.wrapped.Send(ctx, request)
		if errors.Is(err, ErrConnectionRefused) {
			logf("send failed, trying to reconnect")
			err = reconnector.reconnect(ctx)
			if err != nil {
				return "", err
			}
			continue
		}
		return response, err
	}
}

func (reconnector *ReconnectingService) Connect(ctx context.Context) error {
	return reconnector.wrapped.Connect(ctx)
}

func (reconnector *ReconnectingService) reconnect(ctx context.Context) error {
	timeout := reconnector.timeout

	for attempt := 1; attempt <= reconnector.maxAttempts; attempt++ {
		logf("trying to reconnect: attempt #%d", attempt)
		err := reconnector.Connect(ctx)
		if err == nil {
			logf("reconnection succeeded at attempt %d", attempt)
			return nil
		}

		ticker := time.NewTicker(timeout)
		timeout = timeout * time.Duration(reconnector.multiplier)
		select {
		case <-ticker.C:
			ticker.Stop()
		case <-ctx.Done():
			ticker.Stop()
			return ctx.Err()
		}
	}

	return fmt.Errorf("reconnection failed")
}

func main() {
	log.SetFlags(0)

	log.Println("reconnection failed case:")
	doCase(context.Background(), 3)
	log.Println()

	log.Println("reconnection succeeded case:")
	doCase(context.Background(), 5)
	log.Println()

	log.Println("reconnection timeout case:")
	ctx, cancel := context.WithCancel(context.Background())
	time.AfterFunc(time.Millisecond*1000, cancel)
	doCase(ctx, 5)
}

func doCase(ctx context.Context, maxAttempts int) {
	start = time.Now()

	var service Service
	service = &ServiceImplementation{}
	service = &ReconnectingService{
		wrapped:     service,
		maxAttempts: maxAttempts,
		timeout:     100 * time.Millisecond,
		multiplier:  2,
	}

	response, err := service.Send(ctx, "request")
	if err != nil {
		logf("request failed: %v", err)
	} else {
		logf("request succeeded, response received: %s", response)
	}
}

func logf(format string, v ...interface{}) {
	log.Printf(strconv.FormatInt(time.Now().Sub(start).Milliseconds(), 10)+"ms: "+format+"\n", v...)
}
