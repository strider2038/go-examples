package di

type ServiceFactory interface {
	CreateService() Service
}
