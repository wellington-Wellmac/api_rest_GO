package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Item struct {
	ID    int    json:"id"
	Name  string json:"name"
	Price float64 json:"price"
}

var items = []Item{
	{ID: 1, Name: "Item 1", Price: 10.0},
	{ID: 2, Name: "Item 2", Price: 20.0},
	{ID: 3, Name: "Item 3", Price: 30.0},
}

func getItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(items)
}

func getItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*") // Permite CORS
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(items)
}

func main() {
	http.HandleFunc("/api/items", getItems)

	fmt.Println("Servidor iniciado na porta 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}