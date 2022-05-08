package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/pkg/errors"
)

func DoError() error {
	return errors.WithStack(fmt.Errorf("error"))
}

type stackTracer interface {
	StackTrace() errors.StackTrace
}

func main() {
	err := fmt.Errorf("error hanneped: %w", DoError())

	var tracer stackTracer
	if errors.As(err, &tracer) {
		stackTrace := tracer.StackTrace()
		bytes, err := json.Marshal(struct {
			Message    string
			StackTrace errors.StackTrace
		}{
			Message:    err.Error(),
			StackTrace: stackTrace,
		})
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(string(bytes))
	}
}
