package di

type Service interface {
	ProcessMessage(message string) error
}
