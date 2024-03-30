package repository

import (
	"github.com/jmoiron/sqlx"
	"refactor/model"
)

type NasaImage interface {
	CreateNasaImage(sqlQuery string, nasaRes model.NasaImageResponse) error
}

type SalaryCalculation interface {
	CreateSalaryCalculation(sqlQuery string, req model.SalaryRequest, res model.SalaryResponse) error
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
