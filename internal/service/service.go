package service

import (
	"bcc-hackathon-go/pkg/models"
	"bcc-hackathon-go/pkg/modules"
)

type HrCall interface {
	HrCallLogic(jsonInput modules.Request) (code int, any models.HrCall)
}

type TechCall interface {
	TechCallLogic(jsonInput modules.Request) (code int, any models.TechCall)
}

type InfoFromNumber interface {
	InfoFromNumberLogic(jsonInput modules.Request) (code int, any models.InfoFromNumber)
}

type Service struct {
	HrCall
	InfoFromNumber
	TechCall
}

func NewService() *Service {
	return &Service{
		HrCall:         NewHrCallService(),
		InfoFromNumber: NewInfoFromNumberService(),
		TechCall:       NewTechCallService(),
	}
}
