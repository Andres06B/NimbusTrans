package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	dto "proyecto/DTO"
	models "proyecto/Models"
	"proyecto/Services/implements"
	"proyecto/Services/interfaces"
	"strconv"

	"github.com/gorilla/mux"
)

type ConductorController struct {
	l  *log.Logger
	cu interfaces.ConductorService
}

func NewConductorController(l *log.Logger) *ConductorController {
	return &ConductorController{l, &implements.ConductorService{}}
}

func (cs *ConductorController) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	conductor := cs.cu.GetAll()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(conductor)
}

func (cs *ConductorController) GetAllById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	conductor, result := cs.cu.GetById(id)
	if result["message"] == "Obtenido correctamente" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(conductor)
	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(result)
	}
}

func (cs *ConductorController) Create(w http.ResponseWriter, r *http.Request) {
	var conductor dto.ConductorDTO
	err := json.NewDecoder(r.Body).Decode(&conductor)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	data := models.Conductor{
		Id:           conductor.Id,
		Nombre:       conductor.Nombre,
		Cedula:       conductor.Cedula,
		Num_licencia: conductor.NumLicencia,
		Estado:       conductor.Estado,
	}
	result := cs.cu.Create(data)
	if result["message"] == "Creado correctamente" {
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(result)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(result)
	}

}

func (cs *ConductorController) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	var conductor dto.ConductorDTO
	err := json.NewDecoder(r.Body).Decode(&conductor)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	data := models.Conductor{
		Id:           id,
		Nombre:       conductor.Nombre,
		Cedula:       conductor.Cedula,
		Num_licencia: conductor.NumLicencia,
		Estado:       conductor.Estado,
	}
	result := cs.cu.Update(data)
	if result["message"] == "Actualizado correctamente" {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(result)
	}
}

func (cs *ConductorController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	result := cs.cu.Delete(id)
	if result["message"] == "Eliminado correctamente" {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(result)
	}
}
