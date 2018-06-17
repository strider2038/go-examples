package di

type ConcreteServiceFactory struct{}

func (factory ConcreteServiceFactory) CreateService() Service {
	service := &ConcreteService{}

	return service
}
