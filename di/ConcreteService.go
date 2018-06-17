package di

import (
	"fmt"
)

type ConcreteService struct{}

func (service *ConcreteService) ProcessMessage(message string) error {
	fmt.Println("ConcreteService has processed message:", message)

	return nil
}
