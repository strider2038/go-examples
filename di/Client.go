package di

// type is private to prevent creating it without a constructor
type client struct {
	// interface is a pointer type!
	// do not use *ServiceFactory because it is a pointer to pointer
	// if you use struct instead if interface you should use pointer => *ConcreteServiceFactory
	serviceFactory ServiceFactory
}

// equivalent of class constructor
func NewClient(factory ServiceFactory) *client {
	return &client{factory}
}

// equivalent of class method
// for services you should always use pointer receiver
func (client *client) Run() {
	service := client.serviceFactory.CreateService()
	service.ProcessMessage("Client message")
}
