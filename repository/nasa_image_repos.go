package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"refactor/model"
)

type NasaImagePostgres struct {
	db *sqlx.DB
}

func NewNasaImagePostgres(db *sqlx.DB) *NasaImagePostgres {
	return &NasaImagePostgres{
		db: db,
	}
}

func (p *NasaImagePostgres) CreateNasaImage(sqlQuery string, nasaRes model.NasaImageResponse) error {
	fmt.Println("hello create")

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
	log.Println("Inserted data")
	_ = tx.Commit()
	return nil
}

func (p *NasaImagePostgres) GetAllNasaImage(sqlQuery string, res model.AllNasaDataResponse) error {

	tx, err := p.db.Begin()
	if err != nil {
		return err
	}
	data, err := tx.Query(sqlQuery)

	row, err := GetCalcData(data)
	if err != nil {
		return err
	}
	log.Println("row", row)
	//for r := range row {
	//	res.Data = append(res.Data, r)
	//}
	return nil
}
