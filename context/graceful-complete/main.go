package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"
)

const (
	workDuration        = 1000 * time.Millisecond
	completeTimeout     = 1500 * time.Millisecond
	gracefulStopTimeout = 500 * time.Millisecond
)

var start = time.Now()

func longWork(ctx context.Context, data string) (string, error) {
	timer := time.NewTimer(workDuration)

	select {
	case <-timer.C:
		return fmt.Sprintf("data '%s' processed", data), nil
	case <-ctx.Done():
		return "", ctx.Err()
	}
}

func main() {
	log.SetFlags(0)

	ctx, cancel := context.WithTimeout(context.Background(), gracefulStopTimeout)
	defer cancel()

	doWorkWithGracefulComplete(ctx)
}

func doWorkWithGracefulComplete(ctx context.Context) {
	logf("starting work")

	completeContext, cancel := context.WithTimeout(context.Background(), completeTimeout)
	defer cancel()

	done := make(chan struct{}, 1)

	go func() {
		result, err := longWork(completeContext, "input")
		if err != nil {
			logf("work failed: %v", err)
		} else {
			logf("work completed: %s", result)
		}

		done <- struct{}{}
	}()

	select {
	case <-done:
		logf("work is done")
	case <-completeContext.Done():
		logf("work canceled")
	}
}

func logf(format string, v ...interface{}) {
	log.Printf(strconv.FormatInt(time.Now().Sub(start).Milliseconds(), 10)+"ms: "+format+"\n", v...)
}
