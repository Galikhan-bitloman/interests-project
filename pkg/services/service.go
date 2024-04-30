package services

import (
	"interests-project/pkg/model"
	"interests-project/pkg/repository"
)

type Nasa interface {
	GetNasaImageService() (*model.DataNasaResponse, error)
	GetAllNasaImageService(sqlQuery string) (*model.AllNasaDataResponse, error)
}

type SalaryCalculation interface {
	CalculateSalaryService(req model.SalaryRequest) (*model.SalaryResponse, error)
}

type Service struct {
	Nasa
	SalaryCalculation
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Nasa:              NewNasaImageService(repos),
		SalaryCalculation: NewCalculateSalaryService(repos),
	}
}
