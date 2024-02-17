package models


type Vehicle struct {
	Id     int    `json:"id"`
	Brand  string `json:"brand"`
	Model  string `json:"model"`
	Volume int    `json:"volume"`
	Weight int    `json:"weight"`
	CO2    int    `json:"co2"`
}

type VehicleData struct {
	Id     map[string]int    `json:"id"`
	Brand  map[string]string `json:"brand"`
	Model  map[string]string `json:"model"`
	Volume map[string]int    `json:"volume"`
	Weight map[string]int    `json:"weight"`
	CO2    map[string]int    `json:"co2"`
}