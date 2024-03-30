package model

import "time"

type Person struct {
	Name string `json:"name"`
	Age  int8   `json:"age"`
}

type PersonResponse struct {
	NameRes string `json:"nameRes"`
	AgeRes  int8   `json:"ageRes"`
}

type CheckReq struct {
	Body interface{} `json:"body"`
}

type CheckRes struct {
	Res interface{} `json:"res"`
}

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

type ErrorBody struct {
	Msg error `json:"msg"`
}

func (e ErrorBody) Error(err error) string {
	//TODO implement me
	panic("implement me")
}

type NasaImageResponse struct {
	Copyright       string `json:"copyright"`
	Date            string `json:"date"`
	Explanation     string `json:"explanation"`
	Hdurl           string `json:"hdurl"`
	Media_type      string `json:"media_type"`
	Service_version string `json:"service_version"`
	Title           string `json:"title"`
	Url             string `json:"url"`
}

type DataNasaResponse struct {
	Explanation string `json:"explanation"`
	Hdurl       string `json:"hdurl"`
	Title       string `json:"title"`
	Url         string `json:"url"`
	Date        string `json:"date"`
}

type DataResponse struct {
	Data *DataNasaResponse `json:"data"`
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
