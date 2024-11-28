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

type AsignacionController struct {
	l  *log.Logger
	as interfaces.AsignacionService
}

func NewAsignacionController(l *log.Logger) *AsignacionController {
	return &AsignacionController{l, &implements.AsignacionService{}}
}

func (as *AsignacionController) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	asignacion := as.as.GetAll()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(asignacion)
}

func (as *AsignacionController) GetAllById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	asignacion, result := as.as.GetById(idStr)
	if result["message"] == "Obtenido correctamente" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(asignacion)
	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(result)
	}
}

func (as *AsignacionController) Create(w http.ResponseWriter, r *http.Request) {
	var asignacion dto.AsignacionDTO
	err := json.NewDecoder(r.Body).Decode(&asignacion)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	data := models.Asignacion{
		IDconductor: asignacion.Id_conductor,
		IDbus:       asignacion.Id_bus,
		IDruta:      asignacion.Id_ruta,
	}

	result := as.as.Create(data)
	if result["message"] == "Creado correctamente" {
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(result)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(result)
	}
}

func (as *AsignacionController) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	var asignacion dto.AsignacionDTO
	err := json.NewDecoder(r.Body).Decode(&asignacion)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	data := models.Asignacion{
		Id:          id,
		IDconductor: asignacion.Id_conductor,
		IDbus:       asignacion.Id_bus,
		IDruta:      asignacion.Id_ruta,
	}
	result := as.as.Update(data)
	if result["message"] == "Actualizado correctamente" {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(result)
	}
}

func (as *AsignacionController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	result := as.as.Delete(id)
	if result["message"] == "Eliminado correctamente" {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(result)
	}
}
