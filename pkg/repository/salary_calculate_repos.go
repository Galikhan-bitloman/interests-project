package repository

import (
	"interests-project/pkg/model"
	"log"

	"github.com/jmoiron/sqlx"
)

type CalculateSalaryPostgres struct {
	db *sqlx.DB
}

func NewCalculateSalaryPostgres(db *sqlx.DB) *CalculateSalaryPostgres {
	return &CalculateSalaryPostgres{
		db: db,
	}
}

func (c CalculateSalaryPostgres) CreateSalaryCalculationRepository(sqlQuery string, req model.SalaryRequest, res model.SalaryResponse) error {
	tx, err := c.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(sqlQuery, req.Salary, res.NecessarySpending, res.Saving, res.Entertainment)
	if err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			log.Fatal("rollback failed ", rollbackErr)
		}
		return err
	}
	log.Println("Data inserted ")
	_ = tx.Commit()
	return nil
}
