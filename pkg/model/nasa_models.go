package model

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
	Data interface{} `json:"data"`
}

type AllNasaDataResponse struct {
	Data []DataNasaResponse `json:"data"`
}
