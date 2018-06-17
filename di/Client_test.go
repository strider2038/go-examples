package di

import (
	"testing"

	"github.com/stretchr/testify/mock"
)

type mockService struct {
	mock.Mock
}

func (mock *mockService) ProcessMessage(message string) error {
	args := mock.Called(message)

	return args.Error(0)
}

type mockServiceFactory struct {
	mock.Mock
}

func (mock *mockServiceFactory) CreateService() Service {
	args := mock.Called()

	return args.Get(0).(Service)
}

func TestClient_Run_noParameters_messageProcessedByCreatedService(t *testing.T) {
	service := &mockService{}
	serviceFactory := mockServiceFactory{}
	client := NewClient(&serviceFactory)
	serviceFactory.On("CreateService").Return(service)
	service.On("ProcessMessage", "Client message").Return(nil)

	client.Run()

	serviceFactory.AssertExpectations(t)
	service.AssertExpectations(t)
}
