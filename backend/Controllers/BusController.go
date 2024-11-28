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

type BusController struct {
	l  *log.Logger
	bu interfaces.BusService
}

func NewBusController(l *log.Logger) *BusController {
	return &BusController{l, &implements.BusService{}}
}

func (bs *BusController) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	bus := bs.bu.GetAll()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bus)
}

func (bs *BusController) GetAllById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	bus, result := bs.bu.GetById(id)
	if result["message"] == "Obtenido correctamente" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(bus)
	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(result)
	}
}

func (bs *BusController) Create(w http.ResponseWriter, r *http.Request) {
	var bus dto.BusDTO
	err := json.NewDecoder(r.Body).Decode(&bus)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	data := models.Bus{
		Id:        bus.Id,
		Placa:     bus.Placa,
		Modelo:    bus.Modelo,
		Capacidad: bus.Capacidad,
		Estado:    bus.Estado,
	}
	result := bs.bu.Create(data)
	if result["message"] == "Creado correctamente" {
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(result)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(result)
	}

}

func (bs *BusController) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	var bus dto.BusDTO
	err := json.NewDecoder(r.Body).Decode(&bus)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	data := models.Bus{
		Id:        id,
		Placa:     bus.Placa,
		Modelo:    bus.Modelo,
		Capacidad: bus.Capacidad,
		Estado:    bus.Estado,
	}
	result := bs.bu.Update(data)
	if result["message"] == "Actualizado correctamente" {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(result)
	}
}

func (bs *BusController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	result := bs.bu.Delete(id)
	if result["message"] == "Eliminado correctamente" {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(result)
	}
}
