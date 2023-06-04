package main

import (
	"log"
	"net/http"

	"rest-api-go/data"
	"rest-api-go/handlers"

	"github.com/gorilla/mux"
)

func main() {
	// Ler o arquivo de dados
	data.ReadFile()

	// Criação do roteador utilizando o pacote mux
	router := mux.NewRouter().StrictSlash(true)

	// Definição das rotas e seus respectivos manipuladores
	router.HandleFunc("/cars", handlers.ReturnAllCars).Methods("GET")
	router.HandleFunc("/cars/{id}", handlers.ReturnCarsById).Methods("GET")
	router.HandleFunc("/cars/volume/{volume}", handlers.ReturnCarsByVolume).Methods("GET")
	router.HandleFunc("/cars/weight/{weight}", handlers.ReturnCarsByWeight).Methods("GET")
	router.HandleFunc("/cars/co2/{co2}", handlers.ReturnCarsByCO2).Methods("GET")
	router.HandleFunc("/cars/brand/{brand}", handlers.ReturnCarsByBrand).Methods("GET")
	router.HandleFunc("/cars/model/{model}", handlers.ReturnCarsByModel).Methods("GET")
	router.HandleFunc("/cars/{id}", handlers.UpdateCar).Methods("PUT")
	router.HandleFunc("/cars", handlers.CreateCar).Methods("POST")
	router.HandleFunc("/cars/{id}", handlers.RemoveCar).Methods("DELETE")

	// Inicia o servidor na porta 8081
	log.Fatal(http.ListenAndServe(":8081", router))
}
