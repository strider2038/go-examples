package main

import (
	"context"
	"log"

	"github.com/strider2038/go-examples/grpc/messaging"
	"github.com/xeipuuv/gojsonschema"
)

type Messenger struct {
	messaging.UnimplementedMessengerServer
}

var loader = gojsonschema.NewStringLoader(`{
	"type": "object",
	"required": ["name"],
	"properties": {
		"name": {
			"type": "string",
			"minLength": 5
		}
	}
}`)

func NewMessenger() *Messenger {
	return &Messenger{}
}

func (m *Messenger) Dispatch(ctx context.Context, envelope *messaging.Envelope) (*messaging.DispatchResult, error) {
	log.Printf("id: %s, body: %s, properties: %s", envelope.Headers.Id.String(), envelope.Body, envelope.Properties)
	log.Println()

	res, err := gojsonschema.Validate(loader, gojsonschema.NewGoLoader(envelope.Body.AsInterface()))
	if err != nil {
		return nil, err
	}
	if !res.Valid() {
		errs := make([]string, 0)
		for _, resultError := range res.Errors() {
			errs = append(errs, resultError.String())
		}

		return &messaging.DispatchResult{
			Status: messaging.DispatchResult_Fail,
			Errors: errs,
		}, nil
	}

	return &messaging.DispatchResult{
		Status: messaging.DispatchResult_Success,
		Id:     envelope.Headers.Id,
	}, nil
}
