package systemconfig

type ISystemConfigService interface {
	GetSystemConfigList()
}

type systemConfigService struct {
}

func NewSystemConfigService() ISystemConfigService {
	return &systemConfigService{}
}
func (s *systemConfigService) GetSystemConfigList() {
	return
}
