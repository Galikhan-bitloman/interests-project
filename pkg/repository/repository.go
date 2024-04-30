package repository

import (
	"interests-project/model"

	"github.com/jmoiron/sqlx"
)

type NasaImage interface {
	CreateNasaImageRepository(sqlQuery string, nasaRes model.NasaImageResponse) error
	GetAllNasaImageRepository(sqlQuery string) (*[]model.DataNasaResponse, error)
}

type SalaryCalculation interface {
	CreateSalaryCalculationRepository(sqlQuery string, req model.SalaryRequest, res model.SalaryResponse) error
}

type Repository struct {
	NasaImage
	SalaryCalculation
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		NasaImage:         NewNasaImagePostgres(db),
		SalaryCalculation: NewCalculateSalaryPostgres(db),
	}
}
