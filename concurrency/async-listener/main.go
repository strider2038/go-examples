package main

import (
	"context"
	"log"
	"sync"
	"time"
)

func logf(format string, v ...interface{}) {
	log.SetPrefix(time.Now().Format(time.RFC3339))
	log.Printf(format+"\n", v...)
}

func listener(ctx context.Context, input <-chan string, wg *sync.WaitGroup) {
	logf("starting listener...")
	for {
		select {
		case message := <-input:
			logf("receiving message: %s", message)
			time.Sleep(time.Second * 2)
			logf("message '%s' is processed", message)
		case <-ctx.Done():
			logf("listener shut down...")
			wg.Done()
			return
		}
	}
}

func main() {
	ctx, stop := context.WithCancel(context.Background())
	channel := make(chan string)
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go listener(ctx, channel, wg)

	time.Sleep(time.Millisecond)

	messages := []string{"first", "second", "third"}
	for _, message := range messages {
		logf("sending message: %s", message)
		channel <- message
		logf("message '%s' is sent", message)
	}

	stop()
	wg.Wait()
	close(channel)
}
