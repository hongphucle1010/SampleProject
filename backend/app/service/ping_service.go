package service

type IPingService interface {
	Pong() (string, error)
}

type PingService struct {
}

func NewPingService() IPingService {
	return &PingService{}
}

func (s *PingService) Pong() (string, error) {
	return "pong", nil
}