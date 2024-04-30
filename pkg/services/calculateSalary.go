package services

import (
	"errors"
	"interests-project/pkg/model"
	"interests-project/pkg/repository"
	"interests-project/pkg/sqlscripts"
	"log"
)

type CalculateSalary struct {
	repo repository.SalaryCalculation
}

func NewCalculateSalaryService(repo repository.SalaryCalculation) *CalculateSalary {
	return &CalculateSalary{
		repo: repo,
	}
}

func (c *CalculateSalary) CalculateSalaryService(req model.SalaryRequest) (*model.SalaryResponse, error) {
	if req.Salary < 1.0 {
		return nil, errors.New("below zero")
	}

	salaryResp := &model.SalaryResponse{
		Entertainment:     req.Salary * 0.3,
		Saving:            req.Salary * 0.2,
		NecessarySpending: req.Salary * 0.5,
	}
	err := c.repo.CreateSalaryCalculationRepository(sqlscripts.Calc, req, *salaryResp)
	if err != nil {
		log.Fatal(err)
	}
	return salaryResp, nil
}
