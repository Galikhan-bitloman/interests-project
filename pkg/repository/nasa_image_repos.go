package repository

import (
	"interests-project/pkg/model"
	"log"

	"github.com/jmoiron/sqlx"
)

type NasaImagePostgres struct {
	db *sqlx.DB
}

func NewNasaImagePostgres(db *sqlx.DB) *NasaImagePostgres {
	return &NasaImagePostgres{
		db: db,
	}
}

func (p *NasaImagePostgres) CreateNasaImageRepository(sqlQuery string, nasaRes model.NasaImageResponse) error {

	tx, err := p.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(sqlQuery, nasaRes.Url, nasaRes.Explanation, nasaRes.Hdurl, nasaRes.Date, nasaRes.Title)
	if err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			log.Fatal("rollback failed ", rollbackErr)
		}
		return err
	}

	_ = tx.Commit()
	return nil
}

func (p *NasaImagePostgres) GetAllNasaImageRepository(sqlQuery string) (*[]model.DataNasaResponse, error) {
	tx, err := p.db.Begin()
	if err != nil {
		return nil, err
	}
	data, err := tx.Query(sqlQuery)
	if err != nil {
		return nil, err
	}

	row, err := GetCalcData(data)

	if err != nil {
		return nil, err
	}

	return &row, nil
}
