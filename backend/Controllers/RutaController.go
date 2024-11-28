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

type RutaController struct {
	l  *log.Logger
	ru interfaces.RutaService
}

func NewRutaController(l *log.Logger) *RutaController {
	return &RutaController{l, &implements.RutaService{}}
}

func (rs *RutaController) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ruta := rs.ru.GetAll()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ruta)
}

func (rs *RutaController) GetAllById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vr := mux.Vars(r)
	idStr, err := strconv.Atoi(vr["id"])
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	ruta, result := rs.ru.GetById(idStr)
	if result["message"] == "Obtenido correctamente" {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(ruta)
	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(result)
	}
}

func (rs *RutaController) Create(w http.ResponseWriter, r *http.Request) {
	var ruta dto.RutaDTO
	err := json.NewDecoder(r.Body).Decode(&ruta)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	data := models.Ruta{
		Id:           ruta.Id,
		Nombre:       ruta.Nombre,
		Origen:       ruta.Origen,
		Destino:      ruta.Destino,
		Precio_viaje: ruta.PrecioViaje,
		Distancia:    ruta.Distancia,
	}
	result := rs.ru.Create(data)
	if result["message"] == "Creado correctamente" {
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(result)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(result)
	}

}

func (rs *RutaController) Update(w http.ResponseWriter, r *http.Request) {
	vr := mux.Vars(r)
	idStr, err := strconv.Atoi(vr["id"])
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	var ruta dto.RutaDTO
	err = json.NewDecoder(r.Body).Decode(&ruta)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	data := models.Ruta{
		Id:           idStr,
		Nombre:       ruta.Nombre,
		Origen:       ruta.Origen,
		Destino:      ruta.Destino,
		Precio_viaje: ruta.PrecioViaje,
		Distancia:    ruta.Distancia,
	}
	result := rs.ru.Update(data)
	if result["message"] == "Actualizado correctamente" {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(result)
	}
}

func (rs *RutaController) Delete(w http.ResponseWriter, r *http.Request) {
	vr := mux.Vars(r)
	idStr, err := strconv.Atoi(vr["id"])
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	result := rs.ru.Delete(idStr)
	if result["message"] == "Eliminado correctamente" {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(result)
	}
}
