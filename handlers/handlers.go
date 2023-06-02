package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"ProjetoGo/api/data"

	"github.com/gorilla/mux"
)

// ReturnAllCars retorna todos os carros armazenados.
func ReturnAllCars(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data.All_Vehicles)
}

// ReturnCarsByBrand retorna carros filtrados pela marca.
func ReturnCarsByBrand(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	car_brand := params["brand"]

	var cars []data.Vehicle
	for _, car := range data.All_Vehicles {
		if car.Brand == car_brand {
			cars = append(cars, car)
		}
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cars)
}

// ReturnCarsByModel retorna carros filtrados pelo modelo.
func ReturnCarsByModel(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	car_model := vars["model"]

	var cars []data.Vehicle
	for _, car := range data.All_Vehicles {
		if car.Model == car_model {
			cars = append(cars, car)
		}
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cars)
}

// ReturnCarsByVolume retorna carros filtrados pelo volume.
func ReturnCarsByVolume(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	car_volume, err := strconv.Atoi(vars["volume"])
	if err != nil {
		fmt.Println("Unable to convert to string")
	}
	var cars []data.Vehicle
	for _, car := range data.All_Vehicles {
		if car.Volume == car_volume {
			cars = append(cars, car)
		}
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cars)
}

// ReturnCarsByWeight retorna carros filtrados pelo peso.
func ReturnCarsByWeight(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	car_weight, err := strconv.Atoi(vars["weight"])
	if err != nil {
		fmt.Println("Unable to convert to string")
	}
	var cars []data.Vehicle
	for _, car := range data.All_Vehicles {
		if car.Weight == car_weight {
			cars = append(cars, car)
		}
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cars)
}

// ReturnCarsByCO2 retorna carros filtrados pela emissão de CO2.
func ReturnCarsByCO2(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	car_co2, err := strconv.Atoi(vars["co2"])
	if err != nil {
		fmt.Println("Unable to convert to string")
	}
	var cars []data.Vehicle
	for _, car := range data.All_Vehicles {
		if car.CO2 == car_co2 {
			cars = append(cars, car)
		}
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cars)
}

// ReturnCarsById retorna um carro específico com base no ID.
func ReturnCarsById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	car_id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Println("Unable to convert to string")
	}
	var cars []data.Vehicle
	for _, car := range data.All_Vehicles {
		if car.Id == car_id {
			cars = append(cars, car)
		}
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cars)
}

// CreateCar cria um novo carro.
func CreateCar(w http.ResponseWriter, r *http.Request) {
	var newCar data.Vehicle
	json.NewDecoder(r.Body).Decode(&newCar)
	data.All_Vehicles = append(data.All_Vehicles, newCar)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data.All_Vehicles)
}

// UpdateCar atualiza um carro existente.
func UpdateCar(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	car_id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Println("Unable to convert to string")
	}
	var UpdateCar data.Vehicle
	json.NewDecoder(r.Body).Decode(&UpdateCar)
	for k, v := range data.All_Vehicles {
		if v.Id == car_id {
			data.All_Vehicles = append(data.All_Vehicles[:k], data.All_Vehicles[k+1:]...)
			data.All_Vehicles = append(data.All_Vehicles, UpdateCar)
		}
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data.All_Vehicles)
}

// RemoveCar remove um carro existente.
func RemoveCar(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	car_id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Println("Unable to convert to string")
	}
	for k, v := range data.All_Vehicles {
		if v.Id == car_id {
			data.All_Vehicles = append(data.All_Vehicles[:k], data.All_Vehicles[k+1:]...)
		}
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data.All_Vehicles)
}
