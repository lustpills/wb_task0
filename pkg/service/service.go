package service

type Orders interface {
}

type Service struct {
	Orders
}

func NewService() *Service {
	return &Service{}
}
