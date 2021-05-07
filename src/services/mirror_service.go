package services

import "time"

var (
	MirrorService mirrorServiceInterface = &mirrorService{}
)

type mirrorService struct {}


type mirrorServiceInterface interface {
	GetConfig() (*string, error)
	Activate() (*string, error)
	Deactivate() (*string, error)
	GetLogs(date time.Time) (*string, error)
}

func (m *mirrorService) GetConfig() (*string, error) {
	config := "yaml configuration"
	return &config, nil
}

func (m *mirrorService) Activate() (*string, error) {
	resp := "Mirror database activated"
	return &resp, nil
}

func (m *mirrorService) Deactivate() (*string, error) {
	resp := "Mirror database deactivated"
	return &resp, nil
}

func (m *mirrorService) GetLogs(d time.Time) (*string, error) {
	resp := "mirror logs of "
	return &resp, nil
}