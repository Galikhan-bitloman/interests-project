package repository

import (
	"database/sql"
	"fmt"
	"interests-project/pkg/model"
)

func GetCalcData(all *sql.Rows) ([]model.DataNasaResponse, error) {
	var allData []model.DataNasaResponse

	for all.Next() {
		var data model.DataNasaResponse
		if err := all.Scan(&data.Title, &data.Url, &data.Date, &data.Explanation, &data.Hdurl); err != nil {

			fmt.Println("err in ", err)
			return allData, err
		}
		allData = append(allData, data)
	}
	return allData, nil

}
