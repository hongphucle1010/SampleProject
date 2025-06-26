package service

type ITestService interface {
	GetTest() (string, error)
}

type TestService struct{}

func (s *TestService) GetTest() (string, error) {
	return "test", nil
}
