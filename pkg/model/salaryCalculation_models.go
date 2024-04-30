package model

import "time"

type SalaryResponse struct {
	Entertainment     float64 `json:"entertainment"`
	Saving            float64 `json:"saving"`
	NecessarySpending float64 `json:"necessarySpending"`
}

type DataSalaryResponse struct {
	Data *SalaryResponse `json:"data"`
}

type SalaryRequest struct {
	Salary float64 `json:"salary"`
}

type GetAllCalculation struct {
	SalaryGet         float64   `json:"salaryGet"`
	NecessarySpending float64   `json:"necessarySpending"`
	Savings           float64   `json:"savings"`
	Entertainment     float64   `json:"entertainment"`
	Created_at        time.Time `json:"created_at"`
}

type GetAllCalculationData struct {
	Data []GetAllCalculation `json:"data"`
}
