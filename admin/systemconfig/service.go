package systemconfig

type IService interface {
	GetSystemConfigList()
}

type service struct {
}

func NewService() IService {
	return &service{}
}
func (s *service) GetSystemConfigList() {
	return
}
