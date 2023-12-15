package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/edvaldocebol/API-golang.git/models"
)

func List(w http.ResponseWriter, r *http.Request) {
	todos,err := models.GetAll()

	if err!=nil{
		log.Printf("Erro ao obter registro: %v",err)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}