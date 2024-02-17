package data

import (
	"encoding/json"
	"log"
	"os"
	"rest-api-go/pkg/models"
	"strconv"
)



var All_Vehicles []models.Vehicle

func ReadFile() {
	// Abrir o arquivo JSON
	file, err := os.Open("pkg/data/vehicles.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var Vehicle_Data models.VehicleData

	// Decodificar o arquivo JSON para a struct VehicleData
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Vehicle_Data)
	if err != nil {
		log.Fatal(err)
	}

	// Iterar sobre os dados da struct VehicleData e criar objetos Vehicle correspondentes
	for i := 0; i < len(Vehicle_Data.Id); i++ {
		id := Vehicle_Data.Id[strconv.Itoa(i)]
		brand := Vehicle_Data.Brand[strconv.Itoa(i)]
		model := Vehicle_Data.Model[strconv.Itoa(i)]
		volume := Vehicle_Data.Volume[strconv.Itoa(i)]
		weight := Vehicle_Data.Weight[strconv.Itoa(i)]
		co2 := Vehicle_Data.CO2[strconv.Itoa(i)]

		vehicle := models.Vehicle{
			Id:     id,
			Brand:  brand,
			Model:  model,
			Volume: volume,
			Weight: weight,
			CO2:    co2,
		}

		// Adicionar o veículo à lista de todos os veículos
		All_Vehicles = append(All_Vehicles, vehicle)
	}
}
