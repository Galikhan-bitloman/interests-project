package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"refactor/model"
	"refactor/repository"
	"refactor/sqlscripts"
	"time"
)

type NasaImageService struct {
	repo repository.NasaImage
}

func NewNasaImageService(repo repository.NasaImage) *NasaImageService {
	return &NasaImageService{
		repo: repo,
	}
}

func (r *NasaImageService) GetNasaImageService() (*model.DataNasaResponse, error) {
	url := "https://api.nasa.gov/planetary/apod?api_key=SI41eTW8EMaHkrIUwDq7iVdYqzBDT1bx2Bgh5gCo"

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	client := http.Client{}
	response, err := client.Do(request)

	responseBody, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	to := time.Now()
	to.Date()
	var nasaReq model.NasaImageResponse

	err = json.Unmarshal(responseBody, &nasaReq)

	fmt.Println("nasaReq", nasaReq)
	if err != nil {
		return nil, err
	}
	//TODO: new image appears once a day so need to handle and don't add already existing data (check using date)
	err = r.repo.CreateNasaImageRepository(sqlscripts.InsertNasaData, nasaReq)
	if err != nil {
		return nil, err
	}

	res := model.DataNasaResponse{
		Explanation: nasaReq.Explanation,
		Hdurl:       nasaReq.Hdurl,
		Title:       nasaReq.Title,
		Url:         nasaReq.Url,
		Date:        nasaReq.Date,
	}

	return &res, nil
}
