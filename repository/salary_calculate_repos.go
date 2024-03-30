package repository

import (
	"github.com/jmoiron/sqlx"
	"log"
	"refactor/model"
)

type CalculateSalaryPostgres struct {
	db *sqlx.DB
}

func NewCalculateSalaryPostgres(db *sqlx.DB) *CalculateSalaryPostgres {
	return &CalculateSalaryPostgres{
		db: db,
	}
}

func (c CalculateSalaryPostgres) CreateSalaryCalculation(sqlQuery string, req model.SalaryRequest, res model.SalaryResponse) error {
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
