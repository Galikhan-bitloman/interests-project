package services

import (
	"refactor/model"
	"refactor/repository"
)

type Nasa interface {
	GetNasaImage() (*model.DataNasaResponse, error)
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
