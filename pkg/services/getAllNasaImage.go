package services

import (
	"interests-project/pkg/model"
)

func (r *NasaImageService) GetAllNasaImageService(sqlQuery string) (*model.AllNasaDataResponse, error) {
	resp := model.AllNasaDataResponse{}
	data, err := r.repo.GetAllNasaImageRepository(sqlQuery)

	if err != nil {
		return nil, err
	}

	resp.Data = *data
	return &resp, nil

}
